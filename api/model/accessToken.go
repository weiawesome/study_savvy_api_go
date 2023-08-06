package model

import "gorm.io/gorm"

type AccessToken struct {
	gorm.Model
	UserMail string `gorm:"size:254;foreignKey:Mail"`
	Token    string
	AesKey   string
}

func AccessTokenRelate() string {
	return "AccessToken"
}
