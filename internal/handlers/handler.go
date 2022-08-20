package handlers

import (
	configs "bilatung/config"
	"bilatung/internal/models"
	"bilatung/internal/usecase"
)

type handler struct {
	useCase usecase.UseCase
}

type Handlers interface {
	GetHealtcheck() (*models.Health, error)
}

func NewHandler() Handlers {
	return &handler{
		useCase: configs.GetUsecase(),
	}
}
