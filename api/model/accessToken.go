package model

type AccessToken struct {
	UserMail string `gorm:"size:254;foreignKey:Mail"`
	Token    string
	AesKey   string
}
