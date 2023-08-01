package sql

import (
	"errors"
	"gorm.io/gorm"
	"study_savvy_api_go/api/model"
	StatusUtils "study_savvy_api_go/internal/repository/utils"
)

func (r *Repository) CreateFile(obj model.File) error {
	return r.db.Create(obj).Error
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
func (r *Repository) PreLoadReadFile(obj *model.File, preLoad string) error {
	if result := r.db.Preload(preLoad).Find(&obj); result.Error == nil {
		return StatusUtils.ExistSource{Message: "Source is exist"}
	} else if errors.As(result.Error, &gorm.ErrRecordNotFound) {
		return StatusUtils.NotExistSource{Message: "Source is not exist"}
	} else {
		return StatusUtils.DbError{Message: "Db error"}
	}
}
func (r *Repository) UpdateFile(obj model.File) error {
	return r.db.Model(&obj).Updates(obj).Error
}
func (r *Repository) DeleteFile(obj model.File) error {
	return r.db.Delete(&obj).Error
}
