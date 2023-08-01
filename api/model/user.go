package model

type User struct {
	Mail        string `gorm:"primaryKey;size:254;"`
	Gender      string `gorm:"check:gender IN ('male', 'female', 'other')"`
	Salt        []byte
	Password    string
	Name        string
	ApiKey      ApiKey
	AccessToken AccessToken
	File        []File
}
