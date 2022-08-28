package repositories

import (
	"bilatung/internal/models"
	"context"
	"strings"
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
		"tags":       strings.Join(params.Tags, ","),
		"created_at": createdAt,
		"updated_at": updatedAt,
	}).Error; err != nil {
		return err
	}

	return nil
}
