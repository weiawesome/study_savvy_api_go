package files

import (
	"strconv"
	"study_savvy_api_go/api/response/files"
	"study_savvy_api_go/api/utils"
	"study_savvy_api_go/internal/repository/sql"
)

type ServiceFiles struct {
	Repository sql.Repository
}

func (m *ServiceFiles) GetFiles(data string, page int) (files.Files, error) {
	var response files.Files
	pageSize, err := strconv.Atoi(utils.EnvPageSize())
	if err != nil {
		pageSize = 10
	}
	if result, totalPages, err := m.Repository.ReadFileByPage(data, page, pageSize); err == nil {
		var Files []files.File
		for _, file := range result {
			Files = append(Files, file.TranslateToResponseFile())
		}
		return files.Files{Data: Files, CurrentPage: page, TotalPages: totalPages}, nil
	} else {
		return response, err
	}
}
