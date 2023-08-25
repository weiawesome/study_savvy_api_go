package sql

import (
	"errors"
	"gorm.io/gorm"
	"study_savvy_api_go/internal/repository/model"
	StatusUtils "study_savvy_api_go/internal/repository/utils"
)

func (r *Repository) CreateUser(obj model.User) error {
	return r.dbMaster.Create(obj).Error
}
func (r *Repository) ReadUser(obj *model.User) error {
	if result := r.dbSlave.First(&obj); result.Error == nil {
		return StatusUtils.ExistSource{Message: "Resource is exist"}
	} else if errors.As(result.Error, &gorm.ErrRecordNotFound) {
		return StatusUtils.NotExistSource{Message: "Resource is not exist"}
	} else {
		return StatusUtils.DbError{Message: "Db error"}
	}
}
func (r *Repository) PreLoadReadUser(obj *model.User, preLoads []string) error {
	tx := r.dbSlave
	for _, p := range preLoads {
		tx = tx.Preload(p)
	}
	if result := tx.Find(&obj); result.Error == nil {
		return StatusUtils.ExistSource{Message: "Resource is exist"}
	} else if errors.As(result.Error, &gorm.ErrRecordNotFound) {
		return StatusUtils.NotExistSource{Message: "Resource is not exist"}
	} else {
		return StatusUtils.DbError{Message: "Db error"}
	}
}
func (r *Repository) UpdateUser(obj model.User) error {
	return r.dbMaster.Model(&obj).Updates(obj).Error
}
func (r *Repository) DeleteUser(obj model.User) error {
	return r.dbMaster.Delete(&obj).Error
}
