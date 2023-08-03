package files

import (
	"os"
	"strconv"
	"study_savvy_api_go/api/response/files"
	"study_savvy_api_go/internal/repository/sql"
)

type ServiceFilesOcr struct {
	Repository sql.Repository
}

func (m *ServiceFilesOcr) GetFilesOcr(data string, page int) (files.OcrFiles, error) {
	var response files.OcrFiles
	var pageSize int
	pageSize, err := strconv.Atoi(os.Getenv("PAGE_SIZE"))
	if err != nil {
		pageSize = 10
	}
	if result, totalPages, err := m.Repository.ReadFileByPageOcr(data, page, pageSize); err == nil {
		var Files []files.File
		for _, file := range result {
			Files = append(Files, file.TranslateToResponseFile())
		}
		return files.OcrFiles{Data: Files, CurrentPage: page, TotalPages: totalPages}, nil
	} else {
		return response, err
	}
}