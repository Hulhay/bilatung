package handlers

import (
	"bilatung/internal/models"
	"context"
)

func (h *handler) Register(ctx context.Context, params *models.RegisterRequest) error {
	err := h.useCase.Register(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (h *handler) Login(ctx context.Context, params *models.LoginRequest) (*models.Auth, error) {
	res, err := h.useCase.Login(ctx, params)
	if err != nil {
		return nil, err
	}
	return res, nil
}
