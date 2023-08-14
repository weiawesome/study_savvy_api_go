package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ApiKey struct {
	Id       string `gorm:"primaryKey;type:char(36);not null"`
	UserMail string `gorm:"size:254;foreignKey:Mail;not null"`
	Key      string `gorm:"type:text;default:null"`
	AesKey   string `gorm:"type:text;default:null"`
}

func ApiKeyRelate() string {
	return "ApiKey"
}
func (a *ApiKey) BeforeCreate(tx *gorm.DB) (err error) {
	a.Id = uuid.New().String()
	return
}
