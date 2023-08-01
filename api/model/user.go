package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Mail        string `gorm:"primaryKey"`
	Gender      string `gorm:"check:gender IN ('male', 'female', 'other')"`
	Salt        []byte
	Password    string
	Name        string
	ApiKey      ApiKey
	AccessToken AccessToken
	File        []File
}
