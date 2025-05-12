package main

import (
	"auth/app/auth"
	"auth/internal/config"
	"auth/internal/constants"
	"auth/internal/models"
	"auth/internal/router"
	"auth/pkg/database"
	"auth/pkg/httpadapter"
	"auth/pkg/jwt"
	"auth/pkg/logger"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	cfg := config.MustLoad()
	console := logger.Setup(cfg.Env)
	console.Info("App is started", slog.String("env", cfg.Env))
	if cfg.Env != constants.EnvProd {
		console.Debug("Logger working in debug mode")
	} else {
		console.Info("Logger working in info mode")
	}

	db, err := database.InitDatabaseWithSchema(cfg.Database)

	if err != nil {
		log.Fatal(err.Error())
	}

	if err := db.AutoMigrate(
		&models.User{},
	); err != nil {
		console.Info("Migration failed: %v", err)
	}

	r := router.New() // например, Gin Engine

	jwtSrv := jwt.NewJWTService(
		jwt.JWTServiceConfig{
			AccessSecret:  cfg.Jwt.AccessSecret,
			RefreshSecret: cfg.Jwt.RefreshSecret,
			Issuer:        "myapp",
			AccessTTL:     15 * time.Minute,
			RefreshTTL:    7 * 24 * time.Hour,
		})

	authRepo := auth.NewAuthRepository(db)
	authService := auth.NewAuthService(authRepo)

	authController := auth.NewAuthController(authService, jwtSrv)
	authGroup := r.Group("/api/auth")
	authController.RegisterRoutes(authGroup)

	r.Any("/health", httpadapter.WrapWithAdditionalContext(func(ctx context.Context, c *gin.Context) (interface{}, error) {
		return gin.H{"message": "alive"}, nil
	}))

	srv := &http.Server{
		Addr:        cfg.HttpConfig.Address,
		Handler:     r,
		ReadTimeout: cfg.HttpConfig.Timeout,
		IdleTimeout: cfg.HttpConfig.IdleTimeout,
	}

	go func() {
		log.Printf("🚀 Listening on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("⚠️ Shutdown signal received, exiting...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("👋 Server stopped gracefully")

}
