package auth

import (
	"auth/internal/models"
	"auth/pkg/httpadapter"
	"auth/pkg/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	srv AuthService
	jwt *jwt.JWTService
}

func NewAuthController(service AuthService, jwtService *jwt.JWTService) *AuthController {
	return &AuthController{
		srv: service,
		jwt: jwtService,
	}
}

func (c *AuthController) RegisterRoutes(rg *gin.RouterGroup) {

	rg.POST("/register", httpadapter.WrapWithoutAdditionalContext(c.register))
	rg.POST("/login", httpadapter.WrapWithoutAdditionalContext(c.login))
	rg.Any("/validate", httpadapter.WrapWithoutAdditionalContext(c.validate))

}

func (c *AuthController) register(ctx *gin.Context) (interface{}, error) {
	var dto CreateUserDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		return nil, err
	}
	fmt.Println(dto)
	if err := c.srv.Register(ctx.Request.Context(), dto); err != nil {
		return nil, err
	}
	return gin.H{"message": "Created"}, nil
}

func (c *AuthController) _updateTokenPair(ctx *gin.Context, access string, refresh string) {
	ctx.SetCookie("access_token", access, int(c.jwt.AccTTL.Seconds()), "/", "", false, true)
	ctx.SetCookie("refresh_token", refresh, int(c.jwt.RefTTL.Seconds()), "/", "", false, true)
}

func (c *AuthController) login(ctx *gin.Context) (interface{}, error) {

	var creds AuthUserDTO
	if err := ctx.ShouldBindJSON(&creds); err != nil {
		return nil, err
	}
	user, err := c.srv.Authenticate(ctx.Request.Context(), creds)
	if err != nil {
		return nil, err
	}

	access, refresh, err := c.jwt.GenerateTokenPair(ctx, user)

	if err != nil {
		return nil, err
	}

	c._updateTokenPair(ctx, access, refresh)

	return user, nil
}

func (c *AuthController) me(ctx *gin.Context) (interface{}, error) {
	// допустим, AuthMiddleware записал user в контекст
	u, _ := ctx.Request.Context().Value("user").(*models.User)
	return gin.H{"id": u.ID, "username": u.Username, "email": u.Email}, nil
}

func (c *AuthController) validate(ctx *gin.Context) (interface{}, error) {
	claims, err, shouldTokenUpdate := c.jwt.Authenticate(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return gin.H{"message": "Not authorized"}, nil
	}

	// Не получается обновить токен, нужно скорее всего подключать отдельный эндпоинт для этого
	if claims == nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	if shouldTokenUpdate {
		// 5. Обновляем пару токенов
		newAccess, newRefresh, err := c.jwt.GenerateTokenPair(ctx, &models.UserPublic{Id: claims.UserID})
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return gin.H{"message": "Cannot update token"}, nil
		}

		c._updateTokenPair(ctx, newAccess, newRefresh)

		ctx.AbortWithStatus(http.StatusForbidden)
		return gin.H{"message": "Token updated, request again"}, nil
	}
	ctx.Header("x-user-id", claims.UserID)
	ctx.Status(http.StatusOK)

	return gin.H{
		"user_id":  claims.UserID,
		"is_valid": true,
	}, nil
}
