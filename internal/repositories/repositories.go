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

	// User Repository
	CreateUser(ctx context.Context, params *models.RegisterRequest) error
	GetUserByUsername(ctx context.Context, username string) (*models.Users, error)
	GetUserByEmail(ctx context.Context, email string) (*models.Users, error)
}

func NewRepositories(q *gorm.DB) Repositories {
	return &repositories{qry: q}
}
