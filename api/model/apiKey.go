package model

type ApiKey struct {
	UserMail string `gorm:"size:254;foreignKey:Mail"`
	Key      string
	AesKey   string
}
