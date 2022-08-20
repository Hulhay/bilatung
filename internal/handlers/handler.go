package handlers

import (
	configs "bilatung/config"
	"bilatung/internal/usecase"
)

type handler struct {
	useCase usecase.UseCase
}

type Handlers interface {
}

func NewHandler() Handlers {
	return &handler{
		useCase: configs.GetUsecase(),
	}
}
