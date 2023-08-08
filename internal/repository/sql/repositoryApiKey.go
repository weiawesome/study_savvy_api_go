package sql

import (
	"errors"
	"gorm.io/gorm"
	"study_savvy_api_go/api/model"
	StatusUtils "study_savvy_api_go/internal/repository/utils"
)

func (r *Repository) CreateApiKey(obj model.ApiKey) error {
	return r.db.Create(obj).Error
}
func (r *Repository) ReadApiKey(obj *model.ApiKey) error {
	if result := r.db.First(&obj); result.Error == nil {
		return StatusUtils.ExistSource{Message: "Resource is exist"}
	} else if errors.As(result.Error, &gorm.ErrRecordNotFound) {
		return StatusUtils.NotExistSource{Message: "Resource is not exist"}
	} else {
		return StatusUtils.DbError{Message: "Db error"}
	}
}
func (r *Repository) UpdateApiKey(obj model.ApiKey) error {
	return r.db.Model(&obj).Updates(obj).Error
}
func (r *Repository) DeleteApiKey(obj model.ApiKey) error {
	return r.db.Delete(&obj).Error
}

func (r *Repository) PreLoadReadApiKey(obj *model.ApiKey, preLoad string) error {
	if result := r.db.Preload(preLoad).Find(&obj); result.Error == nil {
		return StatusUtils.ExistSource{Message: "Resource is exist"}
	} else if errors.As(result.Error, &gorm.ErrRecordNotFound) {
		return StatusUtils.NotExistSource{Message: "Resource is not exist"}
	} else {
		return StatusUtils.DbError{Message: "Db error"}
	}
}
func (r *Repository) FirstOrCreateApiKey(obj *model.ApiKey) error {
	return r.db.FirstOrCreate(&obj, obj).Error
}
