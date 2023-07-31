package model

import "gorm.io/gorm"

type AccessToken struct {
	gorm.Model
	Mail   string
	Token  string
	AesKey string
}
