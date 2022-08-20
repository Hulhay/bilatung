package usecase

import (
	"bilatung/internal/models"
	"bilatung/internal/shared"
	"context"
	"errors"

	"github.com/google/uuid"
)

func (u *useCase) Register(ctx context.Context, params *models.RegisterRequest) error {

	var (
		err               error
		user              *models.Users
		encryptedPassword string
	)

	// Check username
	user, err = u.repo.GetUserByUsername(ctx, *params.Username)
	if user != nil {
		return errors.New("username sudah digunakan")
	}

	// Check email
	user, err = u.repo.GetUserByEmail(ctx, params.Email.String())
	if user != nil {
		return errors.New("email sudah digunakan")
	}

	token := uuid.New().String()

	if params.UserPhoto == `` {
		params.UserPhoto = shared.USER_PHOTO
	}

	encryptedPassword, err = shared.EncryptPassword(*params.Password)
	if err != nil {
		return err
	}

	req := &models.RegisterRequest{
		Email:     params.Email,
		Password:  &encryptedPassword,
		Token:     token,
		UserPhoto: params.UserPhoto,
		Username:  params.Username,
	}

	err = u.repo.CreateUser(ctx, req)
	if err != nil {
		return err
	}

	user, err = u.repo.GetUserByUsername(ctx, *params.Username)
	if err != nil {
		return nil
	}

	req.UserID = user.UserID
	err = u.repo.CreateAuth(ctx, req)
	if err != nil {
		return nil
	}

	return nil
}
