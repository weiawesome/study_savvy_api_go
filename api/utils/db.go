package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"study_savvy_api_go/api/model"
)

var db *gorm.DB

func migrateModel(db *gorm.DB, model interface{}) error {
	err := db.AutoMigrate(model)
	if err != nil {
		return err
	}
	return nil
}

func InitDB() error {
	user := os.Getenv("SQL_DB_USER")
	pwd := os.Getenv("SQL_DB_PASSWORD")
	port := os.Getenv("SQL_DB_PORT")
	dbName := os.Getenv("SQL_DB_NAME")
	dsn := user + ":" + pwd + "@tcp(localhost:" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	if err := migrateModel(db, &model.User{}); err != nil {
		return err
	}
	if err := migrateModel(db, &model.File{}); err != nil {
		return err
	}
	if err := migrateModel(db, &model.ApiKey{}); err != nil {
		return err
	}
	if err := migrateModel(db, &model.AccessToken{}); err != nil {
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}
