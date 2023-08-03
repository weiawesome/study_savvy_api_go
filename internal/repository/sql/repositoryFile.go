package sql

import (
	"errors"
	"gorm.io/gorm"
	"math"
	"study_savvy_api_go/api/model"
	StatusUtils "study_savvy_api_go/internal/repository/utils"
)

func (r *Repository) CreateFile(obj model.File) error {
	return r.db.Create(obj).Error
}
func (r *Repository) ReadFile(obj *model.File) error {
	if result := r.db.First(&obj); result.Error == nil {
		return StatusUtils.ExistSource{Message: "Resource is exist"}
	} else if errors.As(result.Error, &gorm.ErrRecordNotFound) {
		return StatusUtils.NotExistSource{Message: "Resource is not exist"}
	} else {
		return StatusUtils.DbError{Message: "Db error"}
	}
}
func (r *Repository) PreLoadReadFile(obj *model.File, preLoad string) error {
	if result := r.db.Preload(preLoad).Find(&obj); result.Error == nil {
		return StatusUtils.ExistSource{Message: "Resource is exist"}
	} else if errors.As(result.Error, &gorm.ErrRecordNotFound) {
		return StatusUtils.NotExistSource{Message: "Resource is not exist"}
	} else {
		return StatusUtils.DbError{Message: "Db error"}
	}
}
func (r *Repository) UpdateFile(obj model.File) error {
	return r.db.Model(&obj).Updates(obj).Error
}
func (r *Repository) DeleteFile(obj model.File) error {
	return r.db.Delete(&obj).Error
}

func (r *Repository) ReadFileByPage(mail string, page int, pageSize int) ([]model.File, int, error) {
	var files []model.File
	var totalRecords int64

	if err := r.db.Scopes(FileByMail(mail, &totalRecords), OrderByCreatedAt(), Paginate(page, pageSize)).Count(&totalRecords).Find(&files).Error; err != nil {
		return nil, 0, err
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	return files, totalPages, nil
}
func (r *Repository) ReadFileByPageAsr(mail string, page int, pageSize int) ([]model.File, int, error) {
	var files []model.File
	var totalRecords int64

	if err := r.db.Scopes(FileByMail(mail, &totalRecords), FileByType("ASR"), OrderByCreatedAt(), Paginate(page, pageSize)).Count(&totalRecords).Find(&files).Error; err != nil {
		return nil, 0, err
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	return files, totalPages, nil
}
func (r *Repository) ReadFileByPageOcr(mail string, page int, pageSize int) ([]model.File, int, error) {
	var files []model.File
	var totalRecords int64

	if err := r.db.Scopes(FileByMail(mail, &totalRecords), FileByType("OCR"), OrderByCreatedAt(), Paginate(page, pageSize)).Count(&totalRecords).Find(&files).Error; err != nil {
		return nil, 0, err
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	return files, totalPages, nil
}
func FileByType(Type string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("type = ?", Type)
	}
}
func FileByMail(mail string, totalRecords *int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Model(model.File{}).Where("user_mail = ?", mail).Count(totalRecords)
	}
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
func OrderByCreatedAt() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at DESC")
	}
}
