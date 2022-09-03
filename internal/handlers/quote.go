package handlers

import (
	"bilatung/internal/apis/operations/quote"
	"bilatung/internal/models"
	"context"
)

func (h *handler) CreateQuote(ctx context.Context, params *quote.PostQuoteParams) error {
	err := h.useCase.CreateQuote(ctx, params.Body, params.Token)
	if err != nil {
		return err
	}
	return nil
}

func (h *handler) GetRandomQuote(ctx context.Context) (*models.QuotesResponse, error) {
	res, err := h.useCase.GetRandomQuote(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
