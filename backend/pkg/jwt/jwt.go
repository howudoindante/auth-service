package jwt

import (
	"auth/internal/models"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

type JWTServiceConfig struct {
	AccessSecret  string
	RefreshSecret string
	AccessTTL     time.Duration
	RefreshTTL    time.Duration
	Issuer        string
}

type JWTService struct {
	AccSecret []byte
	RefSecret []byte
	Issuer    string
	AccTTL    time.Duration
	RefTTL    time.Duration
}

func NewJWTService(cfg JWTServiceConfig) *JWTService {
	return &JWTService{
		AccSecret: []byte(cfg.AccessSecret),
		RefSecret: []byte(cfg.RefreshSecret),
		Issuer:    cfg.Issuer,
		AccTTL:    cfg.AccessTTL,
		RefTTL:    cfg.RefreshTTL,
	}
}

type Claims struct {
	UserID string `json:"uid"`
	jwt.RegisteredClaims
}

func (s *JWTService) GenerateTokenPair(_ context.Context, u *models.UserPublic) (access, refresh string, err error) {
	now := time.Now()
	// -- access --
	ac := Claims{
		UserID: fmt.Sprintf("%s", u.Id),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.Issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.AccTTL)),
		},
	}

	access, err = jwt.NewWithClaims(jwt.SigningMethodHS256, ac).SignedString(s.AccSecret)

	if err != nil {
		return "", "", err
	}

	// -- refresh --

	rc := ac

	rc.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(now.Add(s.RefTTL))

	refresh, err = jwt.NewWithClaims(jwt.SigningMethodHS256, rc).SignedString(s.RefSecret)

	if err != nil {
		return "", "", err
	}

	return access, refresh, err
}

func (s *JWTService) parseWithSecret(token string, secret []byte) (*Claims, error) {
	parsed, err := jwt.ParseWithClaims(token, Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})

	if err != nil {
		return nil, err
	}
	cl, ok := parsed.Claims.(*Claims)

	if !ok || !parsed.Valid {
		return nil, errors.New("invalid token")
	}
	return cl, nil
}

func (s *JWTService) ValidateAccessToken(_ context.Context, tokenStr string) (*Claims, error) {
	parser := jwt.NewParser()
	claims := &Claims{}
	token, err := parser.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected signing method")
		}
		return s.AccSecret, nil
	},
	)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func (s *JWTService) ValidateRefreshToken(_ context.Context, tokenStr string) (*Claims, error) {
	parser := jwt.NewParser()
	claims := &Claims{}
	token, err := parser.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected signing method")
		}
		return s.RefSecret, nil
	},
	)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func (svc *JWTService) Authenticate(c *gin.Context) (*Claims, error, bool) {
	// 1. Пробуем получить access token из куки
	accessToken, _ := c.Cookie("access_token")

	// 2. Если нет в куках - проверяем Authorization header
	if accessToken == "" {
		authHeader := c.GetHeader("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			accessToken = strings.TrimPrefix(authHeader, "Bearer ")
		}
	}

	// 3. Валидация access токена
	if accessToken != "" {
		claims, err := svc.ValidateAccessToken(c, accessToken)

		if err == nil {
			return claims, nil, false
		}
	}

	// 4. Если нет токена, то кидаем ошибку

	refreshToken, err := c.Cookie("refresh_token")
	if refreshToken == "" || err != nil {
		return nil, errors.New("Not authorized"), false
	}

	// 3. Валидация refresh токена
	claims, err := svc.ValidateRefreshToken(c, refreshToken)

	if err != nil {
		return nil, err, false
	}

	return claims, nil, true
}
