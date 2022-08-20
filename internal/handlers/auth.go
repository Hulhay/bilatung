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
