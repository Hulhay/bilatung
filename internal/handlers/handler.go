package handlers

import (
	configs "bilatung/config"
	"bilatung/internal/apis/operations/quote"
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
	Login(ctx context.Context, params *models.LoginRequest) (*models.Auth, error)
	Logout(ctx context.Context, token string) error

	// Quote Handler
	CreateQuote(ctx context.Context, params *quote.PostQuoteParams) error
}

func NewHandler() Handlers {
	return &handler{
		useCase: configs.GetUsecase(),
	}
}
