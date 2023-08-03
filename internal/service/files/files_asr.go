package files

import (
	"os"
	"strconv"
	"study_savvy_api_go/api/response/files"
	"study_savvy_api_go/internal/repository/sql"
)

type ServiceFilesAsr struct {
	Repository sql.Repository
}

func (m *ServiceFilesAsr) GetFilesAsr(data string, page int) (files.AsrFiles, error) {
	var response files.AsrFiles
	var pageSize int
	pageSize, err := strconv.Atoi(os.Getenv("PAGE_SIZE"))
	if err != nil {
		pageSize = 10
	}
	if result, totalPages, err := m.Repository.ReadFileByPageAsr(data, page, pageSize); err == nil {
		var Files []files.File
		for _, file := range result {
			Files = append(Files, file.TranslateToResponseFile())
		}
		return files.AsrFiles{Data: Files, CurrentPage: page, TotalPages: totalPages}, nil
	} else {
		return response, err
	}
}
