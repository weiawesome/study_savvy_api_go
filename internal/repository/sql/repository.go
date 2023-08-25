package sql

import (
	"gorm.io/gorm"
	"study_savvy_api_go/api/utils"
)

type Repository struct {
	dbMaster *gorm.DB
	dbSlave  *gorm.DB
}

func NewRepository() *Repository {
	return &Repository{dbMaster: utils.GetDBMaster(), dbSlave: utils.GetDBSalve()}
}
