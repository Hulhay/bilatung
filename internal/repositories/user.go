package repositories

import (
	"bilatung/internal/models"
	"context"
	"time"
)

func (r *repositories) CreateUser(ctx context.Context, params *models.RegisterRequest) error {
	var user *models.Users

	createdAt := time.Now()
	updatedAt := createdAt

	tx := r.qry.Begin()
	defer tx.Commit()

	if err := tx.Model(&user).Create(map[string]interface{}{
		"email":      params.Email,
		"password":   params.Password,
		"token":      params.Token,
		"user_photo": params.UserPhoto,
		"username":   params.Username,
		"created_at": createdAt,
		"updated_at": updatedAt,
	}).Error; err != nil {
		tx.Rollback()
		return nil
	}

	return nil
}

func (r *repositories) GetUserByUsername(ctx context.Context, username string) (*models.Users, error) {
	var user *models.Users

	tx := r.qry.Begin()
	defer tx.Commit()

	if err := tx.Model(&user).Where("username = ?", username).First(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return user, nil
}

func (r *repositories) GetUserByEmail(ctx context.Context, email string) (*models.Users, error) {
	var user *models.Users

	tx := r.qry.Begin()
	defer tx.Commit()

	if err := tx.Model(&user).Where("email = ?", email).First(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return user, nil
}

func (r *repositories) GetUserByID(ctx context.Context, userID int64) (*models.Users, error) {
	var user *models.Users

	tx := r.qry.Begin()
	defer tx.Commit()

	if err := tx.Model(&user).Where("user_id = ?", userID).First(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return user, nil
}
