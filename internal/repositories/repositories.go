package repositories

import (
	"gorm.io/gorm"
)

type repositories struct {
	qry *gorm.DB
}

type Repositories interface {
}

func NewRepositories(q *gorm.DB) Repositories {
	return &repositories{qry: q}
}
