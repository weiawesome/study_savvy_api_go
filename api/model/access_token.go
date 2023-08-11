package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccessToken struct {
	Id       string `gorm:"primaryKey;type:char(36);not null"`
	UserMail string `gorm:"size:254;foreignKey:Mail;not null"`
	Token    string `gorm:"type:text;default:null"`
	AesKey   string `gorm:"type:text;default:null"`
}

func AccessTokenRelate() string {
	return "AccessToken"
}

func (a *AccessToken) BeforeCreate(tx *gorm.DB) (err error) {
	a.Id = uuid.New().String()
	return
}
