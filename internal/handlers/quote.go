package handlers

import (
	"bilatung/internal/apis/operations/quote"
	"context"
)

func (h *handler) CreateQuote(ctx context.Context, params *quote.PostQuoteParams) error {
	err := h.useCase.CreateQuote(ctx, params.Body, params.Token)
	if err != nil {
		return err
	}
	return nil
}
