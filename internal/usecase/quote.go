package usecase

import (
	"bilatung/internal/models"
	"bilatung/internal/shared"
	"context"
	"errors"
	"strconv"
	"strings"
)

func (u *useCase) CreateQuote(ctx context.Context, params *models.Quotes, token string) error {

	var (
		err    error
		auth   *models.Auth
		rating *models.Ratings
	)

	if strings.TrimSpace(*params.Quote) == "" {
		return errors.New("tidak boleh hanya mengandung white space")
	}

	_, err = strconv.Atoi(*params.Quote)
	if err == nil {
		return errors.New("tidak boleh hanya mengandung angka")
	}

	filteredText := ``
	filteredText, err = shared.RemovePunctuation(*params.Quote)
	if err != nil || len(filteredText) == 0 {
		return errors.New("tidak boleh hanya mengandung tanda baca")
	}

	auth, err = u.repo.GetAuthByToken(ctx, token)
	if err != nil {
		return err
	}

	rating, err = u.repo.GetRatingByUserID(ctx, int(auth.UserID))
	if err != nil {
		return err
	}
	currentQuoteCount := rating.QuoteCount

	req := &models.Quotes{
		Author: params.Author,
		Quote:  params.Quote,
		Tags:   params.Tags,
		UserID: auth.UserID,
	}
	err = u.repo.CreateQuote(ctx, req)
	if err != nil {
		return err
	}

	currentQuoteCount++
	err = u.repo.AddRatingQuoteCount(ctx, int(auth.UserID), int(currentQuoteCount))
	if err != nil {
		return err
	}

	return nil
}
