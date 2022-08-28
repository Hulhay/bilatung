package repositories

import (
	"bilatung/internal/models"
	"context"
)

func (r *repositories) CreateRating(ctx context.Context, params *models.Ratings) error {
	var rating *models.Ratings

	tx := r.qry.Begin()
	defer tx.Commit()

	if err := tx.Model(&rating).Create(map[string]interface{}{
		"user_id": params.UserID,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *repositories) GetRatingByUserID(ctx context.Context, userID int) (*models.Ratings, error) {
	var rating *models.Ratings

	tx := r.qry.Begin()
	defer tx.Commit()

	if err := tx.Model(&rating).Where("user_id = ?", userID).First(&rating).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return rating, nil
}

func (r *repositories) AddRatingQuoteCount(ctx context.Context, userID int, quoteCount int) error {
	var rating *models.Ratings

	tx := r.qry.Begin()
	defer tx.Commit()

	if err := tx.Model(&rating).Where("user_id = ?", userID).Updates(map[string]interface{}{
		"quote_count": quoteCount,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
