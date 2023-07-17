package ai_predict

import (
	"errors"
	"fmt"
	"mime/multipart"
	"study_savvy_api_go/api/request/ai_predict/utils"
)

type MultipartRequestAsr struct {
	File   *multipart.FileHeader `json:"file"`
	Prompt string                `json:"prompt"`
}

func (r *MultipartRequestAsr) Validate() error {
	if r.File == nil {
		return fmt.Errorf("file is required")
	}
	allowedExtensions := map[string]bool{
		".mp3": true,
		".wav": true,
		".m4a": true,
	}

	ext := utils.GetFileExtension(r.File.Filename)
	if !allowedExtensions[ext] {
		return errors.New("unsupported file format")
	}

	return nil
}
