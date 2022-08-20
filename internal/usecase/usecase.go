package usecase

import (
	"bilatung/internal/models"
	"bilatung/internal/repositories"
	"context"
)

type useCase struct {
	repo repositories.Repositories
}

type UseCase interface {
	Register(ctx context.Context, params *models.RegisterRequest) error
	Login(ctx context.Context, params *models.LoginRequest) (*models.Auth, error)
	Logout(ctx context.Context, token string) error
}

func NewUseCase(r repositories.Repositories) UseCase {
	return &useCase{repo: r}
}
