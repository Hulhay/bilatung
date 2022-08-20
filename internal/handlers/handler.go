package handlers

import (
	configs "bilatung/config"
	"bilatung/internal/models"
	"bilatung/internal/usecase"
	"context"
)

type handler struct {
	useCase usecase.UseCase
}

type Handlers interface {
	GetHealtcheck() (*models.Health, error)

	// Auth Handler
	Register(ctx context.Context, params *models.RegisterRequest) error
}

func NewHandler() Handlers {
	return &handler{
		useCase: configs.GetUsecase(),
	}
}
