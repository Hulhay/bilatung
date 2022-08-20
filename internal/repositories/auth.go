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

func (r *repositories) GetAuthByUsername(ctx context.Context, username string) (*models.Auth, error) {
	var auth *models.Auth

	tx := r.qry.Begin()
	defer tx.Commit()

	if err := tx.Model(&auth).Where("username = ?", username).First(&auth).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return auth, nil
}

func (r *repositories) Login(ctx context.Context, params *models.LoginRequest) error {
	var auth *models.Auth

	tx := r.qry.Begin()
	defer tx.Commit()

	if err := tx.Model(&auth).Where("username = ?", params.Username).Updates(map[string]interface{}{
		"islogin": true,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *repositories) GetAuthByToken(ctx context.Context, token string) (*models.Auth, error) {
	var auth *models.Auth

	tx := r.qry.Begin()
	defer tx.Commit()

	if err := tx.Model(&auth).Where("token = ?", token).First(&auth).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return auth, nil
}

func (r *repositories) Logout(ctx context.Context, token string) error {
	var auth *models.Auth

	tx := r.qry.Begin()
	defer tx.Commit()

	if err := tx.Model(&auth).Where("token = ?", token).Updates(map[string]interface{}{
		"islogin": false,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
