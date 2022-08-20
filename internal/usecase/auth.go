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

func (u *useCase) Login(ctx context.Context, params *models.LoginRequest) (*models.Auth, error) {

	var (
		err  error
		user *models.Users
		auth *models.Auth
	)

	user, err = u.repo.GetUserByUsername(ctx, *params.Username)

	// Check username
	if err != nil && user == nil {
		return nil, errors.New("username not found")
	}

	// Check password
	err = shared.CheckPassword(*params.Password, user.Password)
	if err != nil {
		return nil, errors.New("wrong password")
	}

	auth, err = u.repo.GetAuthByUsername(ctx, *params.Username)

	// Check isLogin
	if auth.Islogin == true {
		return nil, errors.New("you have logged in")
	}

	err = u.repo.Login(ctx, params)
	if err != nil {
		return nil, err
	}

	res := &models.Auth{
		Token: auth.Token,
	}

	return res, nil
}
