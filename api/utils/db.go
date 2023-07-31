package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

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
	return nil
}

func GetDB() *gorm.DB {
	return db
}
