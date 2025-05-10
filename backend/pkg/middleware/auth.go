package middleware

import (
	"auth/internal/models"
	"auth/pkg/jwt"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

const ContextUserKey = "user"

// AuthMiddleware проверяет Bearer‑токен и сохраняет User в контексте
func AuthMiddleware(jwtSvc *jwt.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Вытаскиваем access-токен из cookie
		accessToken, err := c.Cookie("access_token")
		var claims *jwt.Claims

		if err == nil {
			claims, err = jwtSvc.ValidateAccessToken(c.Request.Context(), accessToken)
		}

		// 2. Если access валиден — кладём user и идём дальше
		if err == nil {
			ctx := context.WithValue(c.Request.Context(), ContextUserKey, claims.UserID)
			c.Request = c.Request.WithContext(ctx)
			c.Next()
			return
		}

		refToken, err := c.Cookie("refresh_token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "refresh token missing"})
			return
		}
		claims, err = jwtSvc.ValidateRefreshToken(c.Request.Context(), refToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
			return
		}
		newAccess, newRefresh, err := jwtSvc.GenerateTokenPair(c.Request.Context(), &models.UserPublic{Id: claims.UserID})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "could not refresh tokens"})
			return
		}
		c.SetCookie("access_token", newAccess, int(jwtSvc.AccTTL.Seconds()), "/", "", false, true)
		c.SetCookie("refresh_token", newRefresh, int(jwtSvc.RefTTL.Seconds()), "/", "", false, true)

		// Кладём user в контекст и идём дальше
		ctx := context.WithValue(c.Request.Context(), ContextUserKey, claims.UserID)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
		return

	}
}
