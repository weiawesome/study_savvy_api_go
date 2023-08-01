package sql

import (
	"gorm.io/gorm"
	"study_savvy_api_go/api/utils"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository() *Repository {
	return &Repository{db: utils.GetDB()}
}
