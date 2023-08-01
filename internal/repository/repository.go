package repository

import (
	"errors"
	"gorm.io/gorm"
	"study_savvy_api_go/api/model"
	"study_savvy_api_go/api/utils"
	StatusUtils "study_savvy_api_go/internal/repository/utils"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository() *Repository {
	return &Repository{db: utils.GetDB()}
}

func (r *Repository) CreateUser(obj model.User) error {
	return r.db.Create(obj).Error
}
func (r *Repository) CreateFile(obj model.File) error {
	return r.db.Create(obj).Error
}
func (r *Repository) CreateApiKey(obj model.ApiKey) error {
	return r.db.Create(obj).Error
}
func (r *Repository) CreateAccessToken(obj model.AccessToken) error {
	return r.db.Create(obj).Error
}

func (r *Repository) ReadUser(obj *model.User) error {
	if result := r.db.First(&obj); result.Error == nil {
		return StatusUtils.ExistSource{Message: "Source is exist"}
	} else if errors.As(result.Error, &gorm.ErrRecordNotFound) {
		return StatusUtils.NotExistSource{Message: "Source is not exist"}
	} else {
		return StatusUtils.DbError{Message: "Db error"}
	}
}
func (r *Repository) ReadFile(obj *model.File) error {
	if result := r.db.First(&obj); result.Error == nil {
		return StatusUtils.ExistSource{Message: "Source is exist"}
	} else if errors.As(result.Error, &gorm.ErrRecordNotFound) {
		return StatusUtils.NotExistSource{Message: "Source is not exist"}
	} else {
		return StatusUtils.DbError{Message: "Db error"}
	}
}
func (r *Repository) ReadApiKey(obj *model.ApiKey) error {
	if result := r.db.First(&obj); result.Error == nil {
		return StatusUtils.ExistSource{Message: "Source is exist"}
	} else if errors.As(result.Error, &gorm.ErrRecordNotFound) {
		return StatusUtils.NotExistSource{Message: "Source is not exist"}
	} else {
		return StatusUtils.DbError{Message: "Db error"}
	}
}
func (r *Repository) ReadAccessToken(obj *model.AccessToken) error {
	if result := r.db.First(&obj); result.Error == nil {
		return StatusUtils.ExistSource{Message: "Source is exist"}
	} else if errors.As(result.Error, &gorm.ErrRecordNotFound) {
		return StatusUtils.NotExistSource{Message: "Source is not exist"}
	} else {
		return StatusUtils.DbError{Message: "Db error"}
	}
}

func (r *Repository) PreLoadRead(obj *interface{}, preLoad string) error {
	if result := r.db.Preload(preLoad).Find(&obj); result.Error == nil {
		return StatusUtils.ExistSource{Message: "Source is exist"}
	} else if errors.As(result.Error, &gorm.ErrRecordNotFound) {
		return StatusUtils.NotExistSource{Message: "Source is not exist"}
	} else {
		return StatusUtils.DbError{Message: "Db error"}
	}
}
func (r *Repository) Update(obj interface{}) error {
	return r.db.Model(&obj).Updates(obj).Error
}
func (r *Repository) Delete(obj interface{}) error {
	return r.db.Delete(&obj).Error
}
