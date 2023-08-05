package ai_predict

import (
	"errors"
	"fmt"
	"mime/multipart"
	utils "study_savvy_api_go/api/request/ai_predict/utils"
)

type Ocr struct {
	File   *multipart.FileHeader `json:"files"`
	Prompt string                `json:"prompt"`
}

func (r *Ocr) Validate() (string, error) {
	if r.File == nil {
		return "", fmt.Errorf("files is required")
	}
	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	ext := utils.GetFileExtension(r.File.Filename)
	if !allowedExtensions[ext] {
		return "", errors.New("unsupported files format")
	}

	return ext, nil
}
