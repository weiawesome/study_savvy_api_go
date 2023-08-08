package model

import (
	"study_savvy_api_go/api/response/files"
	"time"
)

type File struct {
	UserMail  string                 `gorm:"size:254;foreignKey:Mail;not null"`
	Id        string                 `gorm:"primaryKey;type:char(36)"`
	Status    string                 `gorm:"type:enum('SUCCESS','FAILURE','PENDING');not null"`
	Type      string                 `gorm:"type:enum('OCR','ASR');not null"`
	CreatedAt time.Time              `gorm:"not null"`
	Resource  string                 `gorm:"not null"`
	Result    map[string]interface{} `gorm:"type:json;default:null"`
}

func FileRelate() string {
	return "File"
}

func (f File) TranslateToResponseFile() files.File {
	return files.File{
		FileId:   f.Id,
		FileType: f.Type,
		FileTime: f.CreatedAt,
		Status:   f.Status,
	}
}
