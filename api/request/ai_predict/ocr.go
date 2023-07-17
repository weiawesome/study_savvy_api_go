package ai_predict

import (
	"errors"
	"fmt"
	"mime/multipart"
	utils "study_savvy_api_go/api/request/ai_predict/utils"
)

type MultipartRequestOcr struct {
	File   *multipart.FileHeader `json:"file"`
	Prompt string                `json:"prompt"`
}

func (r *MultipartRequestOcr) Validate() error {
	if r.File == nil {
		return fmt.Errorf("file is required")
	}
	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	ext := utils.GetFileExtension(r.File.Filename)
	if !allowedExtensions[ext] {
		return errors.New("unsupported file format")
	}

	return nil
}
