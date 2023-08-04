package model

import (
	"study_savvy_api_go/api/response/files"
	"time"
)

type File struct {
	UserMail  string `gorm:"size:254;foreignKey:Mail"`
	Id        string `gorm:"primaryKey;type:char(36)"`
	Status    string `gorm:"check:status IN ('PENDING','SUCCESS','FAILURE')"`
	Type      string `gorm:"check:type IN ('ASR','OCR')"`
	CreatedAt time.Time
	Resource  string
	Result    files.SpecificFile `gorm:"type:json"`
}

func FileRelateName() string {
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
