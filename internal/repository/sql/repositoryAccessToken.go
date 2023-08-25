package sql

import (
	"errors"
	"gorm.io/gorm"
	"study_savvy_api_go/internal/repository/model"
	StatusUtils "study_savvy_api_go/internal/repository/utils"
)

func (r *Repository) CreateAccessToken(obj model.AccessToken) error {
	return r.dbMaster.Create(obj).Error
}
func (r *Repository) ReadAccessToken(obj *model.AccessToken) error {
	if result := r.dbSlave.First(&obj); result.Error == nil {
		return StatusUtils.ExistSource{Message: "Resource is exist"}
	} else if errors.As(result.Error, &gorm.ErrRecordNotFound) {
		return StatusUtils.NotExistSource{Message: "Resource is not exist"}
	} else {
		return StatusUtils.DbError{Message: "Db error"}
	}
}
func (r *Repository) UpdateAccessToken(obj model.AccessToken) error {
	return r.dbMaster.Model(&obj).Updates(obj).Error
}
func (r *Repository) DeleteAccessToken(obj model.AccessToken) error {
	return r.dbMaster.Delete(&obj).Error
}

func (r *Repository) PreLoadReadAccessToken(obj *model.AccessToken, preLoad string) error {
	if result := r.dbSlave.Preload(preLoad).Find(&obj); result.Error == nil {
		return StatusUtils.ExistSource{Message: "Resource is exist"}
	} else if errors.As(result.Error, &gorm.ErrRecordNotFound) {
		return StatusUtils.NotExistSource{Message: "Resource is not exist"}
	} else {
		return StatusUtils.DbError{Message: "Db error"}
	}
}
func (r *Repository) FirstOrCreateAccessToken(obj *model.AccessToken) error {
	return r.dbMaster.FirstOrCreate(&obj, obj).Error
}
