package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"study_savvy_api_go/api/model"
)

var db *gorm.DB

func connectDB() (*gorm.DB, error) {
	user := os.Getenv("SQL_DB_USER")
	pwd := os.Getenv("SQL_DB_PASSWORD")
	port := os.Getenv("SQL_DB_PORT")
	dbName := os.Getenv("SQL_DB_NAME")
	dsn := user + ":" + pwd + "@tcp(localhost:" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Print(dsn)
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
