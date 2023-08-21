package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"study_savvy_api_go/internal/repository/model"
)

var db *gorm.DB

func connectDB() (*gorm.DB, error) {
	user := EnvMySqlUser()
	password := EnvMySqlPassword()
	address := EnvMySqlAddress()
	dbName := EnvMySqlDb()

	dsn := user + ":" + password + "@tcp(" + address + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func InitDB() error {
	var err error
	if db, err = connectDB(); err != nil {
		return err
	}
	if err := db.AutoMigrate(&model.User{}, &model.AccessToken{}, &model.ApiKey{}, &model.File{}); err != nil {
		return err
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() error {
	if db == nil {
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Close()
	if err != nil {
		return err
	}

	return nil
}
