package repositories

import (
	"bilatung/internal/models"
	"context"
	"time"
)

func (r *repositories) CreateQuote(ctx context.Context, params *models.Quotes) error {
	var quote *models.Quotes

	createdAt := time.Now()
	updatedAt := createdAt

	tx := r.qry.Begin()
	defer tx.Commit()

	if err := tx.Model(&quote).Create(map[string]interface{}{
		"author":     params.Author,
		"quote":      params.Quote,
		"user_id":    params.UserID,
		"tags":       params.Tags,
		"created_at": createdAt,
		"updated_at": updatedAt,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (r *repositories) CountQuote(ctx context.Context) (int64, error) {
	var (
		quote *models.Quotes
		count int64
	)

	tx := r.qry.Begin()
	defer tx.Commit()

	if err := tx.Model(&quote).Count(&count).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	return count, nil
}

func (r *repositories) GetQuoteByID(ctx context.Context, quoteID int) (*models.Quotes, error) {
	var quote *models.Quotes

	tx := r.qry.Begin()
	defer tx.Commit()

	if err := tx.Model(&quote).Where("quote_id = ?", quoteID).First(&quote).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return quote, nil
}
