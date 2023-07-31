package model

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type File struct {
	gorm.Model
	Id        uuid.UUID `gorm:"primaryKey;type:char(36)"`
	Status    string    `gorm:"check:status IN ('PENDING','SUCCESS','FAILURE')"`
	Type      string    `gorm:"check:type IN ('ASR','OCR')"`
	CreatedAt time.Time
	Source    string
	Result    datatypes.JSON
}
