package repositories

import (
	"bilatung/internal/models"
	"context"
)

func (r *repositories) CreateAuth(ctx context.Context, params *models.RegisterRequest) error {
	var auth *models.Auth

	tx := r.qry.Begin()
	defer tx.Commit()

	if err := tx.Model(&auth).Create(map[string]interface{}{
		"islogin":  0,
		"token":    params.Token,
		"user_id":  params.UserID,
		"username": params.Username,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
