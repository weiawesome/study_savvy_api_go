package model

import "gorm.io/gorm"

type ApiKey struct {
	gorm.Model
	UserMail string `gorm:"size:254;foreignKey:Mail"`
	Key      string
	AesKey   string
}

func ApiKeyRelate() string {
	return "ApiKey"
}
