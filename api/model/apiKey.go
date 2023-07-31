package model

import "gorm.io/gorm"

type ApiKey struct {
	gorm.Model
	Mail   string
	Key    string
	AesKey string
}
