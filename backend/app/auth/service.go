package auth

import (
	"auth/internal/models"
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, dto CreateUserDTO) error
	Authenticate(ctx context.Context, dto AuthUserDTO) (*models.UserPublic, error)
	Logout(ctx context.Context) error
	Me(ctx context.Context) (*models.UserPublic, error)
}

type authService struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) AuthService {
	return &authService{
		repo: repo,
	}
}

func (svc *authService) Register(ctx context.Context, dto CreateUserDTO) error {
	if existing, _ := svc.repo.GetByUsername(ctx, dto.Username); existing != nil {
		return errors.New("username already taken")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Username: dto.Username,
		Password: string(hash),
		Email:    dto.Email,
	}

	fmt.Println(user)

	if err := svc.repo.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (svc *authService) Authenticate(ctx context.Context, dto AuthUserDTO) (*models.UserPublic, error) {
	user, err := svc.repo.GetByUsername(ctx, dto.Username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
		return nil, err
	}

	return &models.UserPublic{Id: fmt.Sprintf("%s", user.ID), Username: user.Username, Email: user.Email}, nil
}

func (svc *authService) Logout(ctx context.Context) error {
	return nil
}

func (svc *authService) Me(ctx context.Context) (*models.UserPublic, error) {
	return nil, nil
}
