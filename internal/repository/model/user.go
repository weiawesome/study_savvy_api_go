package model

type User struct {
	Mail        string `gorm:"primaryKey;size:254;"`
	Gender      string `gorm:"type:enum('male','female','other');default:'other';not null"`
	Salt        []byte `gorm:"type:Binary(16);default:null"`
	Password    string
	Name        string `gorm:"size:50;not null"`
	ApiKey      ApiKey
	AccessToken AccessToken
	File        []File
}
