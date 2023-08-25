package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"study_savvy_api_go/internal/repository/model"
)

var dbMaster *gorm.DB
var dbSlave *gorm.DB

func connectDBMaster() (*gorm.DB, error) {
	user := EnvMySqlUser()
	password := EnvMySqlPassword()
	address := EnvMySqlMasterAddress()
	dbName := EnvMySqlDb()

	dsn := user + ":" + password + "@tcp(" + address + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func connectDBSlave() (*gorm.DB, error) {
	user := EnvMySqlUser()
	password := EnvMySqlPassword()
	address := EnvMySqlSlaveAddress()
	dbName := EnvMySqlDb()

	dsn := user + ":" + password + "@tcp(" + address + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func InitDB() error {
	var err error
	if dbMaster, err = connectDBMaster(); err != nil {
		return err
	}
	if err := dbMaster.AutoMigrate(&model.User{}, &model.AccessToken{}, &model.ApiKey{}, &model.File{}); err != nil {
		return err
	}
	if dbSlave, err = connectDBSlave(); err != nil {
		return err
	}
	return nil
}

func GetDBMaster() *gorm.DB {
	return dbMaster
}
func GetDBSalve() *gorm.DB {
	return dbSlave
}

func CloseDB() error {
	if dbMaster == nil {
		return nil
	}

	sqlDB, err := dbMaster.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Close()
	if err != nil {
		return err
	}

	return nil
}
