package sql

import (
	"errors"
	"gorm.io/gorm"
	"study_savvy_api_go/api/model"
	StatusUtils "study_savvy_api_go/internal/repository/utils"
)

func (r *Repository) CreateUser(obj model.User) error {
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
func (r *Repository) PreLoadReadUser(obj *model.User, preLoad string) error {
	if result := r.db.Preload(preLoad).Find(&obj); result.Error == nil {
		return StatusUtils.ExistSource{Message: "Source is exist"}
	} else if errors.As(result.Error, &gorm.ErrRecordNotFound) {
		return StatusUtils.NotExistSource{Message: "Source is not exist"}
	} else {
		return StatusUtils.DbError{Message: "Db error"}
	}
}
func (r *Repository) UpdateUser(obj model.User) error {
	return r.db.Model(&obj).Updates(obj).Error
}
func (r *Repository) DeleteUser(obj model.User) error {
	return r.db.Delete(&obj).Error
}
