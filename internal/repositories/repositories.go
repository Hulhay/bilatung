package repositories

import (
	"bilatung/internal/models"
	"context"

	"gorm.io/gorm"
)

type repositories struct {
	qry *gorm.DB
}

type Repositories interface {
	// Auth Repository
	CreateAuth(ctx context.Context, params *models.RegisterRequest) error
	GetAuthByUsername(ctx context.Context, username string) (*models.Auth, error)
	Login(ctx context.Context, params *models.LoginRequest) error
	GetAuthByToken(ctx context.Context, token string) (*models.Auth, error)
	Logout(ctx context.Context, token string) error

	// User Repository
	CreateUser(ctx context.Context, params *models.RegisterRequest) error
	GetUserByUsername(ctx context.Context, username string) (*models.Users, error)
	GetUserByEmail(ctx context.Context, email string) (*models.Users, error)
}

func NewRepositories(q *gorm.DB) Repositories {
	return &repositories{qry: q}
}
