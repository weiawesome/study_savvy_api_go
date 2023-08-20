package files

import (
	"errors"
	"study_savvy_api_go/api/response/files"
	"study_savvy_api_go/internal/repository/model"
	"study_savvy_api_go/internal/repository/sql"
	StatusUtils "study_savvy_api_go/internal/repository/utils"
)

type ServiceFilesSpecific struct {
	Repository sql.Repository
}

func (m *ServiceFilesSpecific) GetFile(data string, id string) (files.SpecificFile, error) {
	var response files.SpecificFile
	File := model.File{Id: id}
	if err := m.Repository.ReadFile(&File); errors.As(err, &StatusUtils.ExistSource{}) {
		if File.UserMail != data {
			return response, StatusUtils.NotExistSource{}
		} else {
			return File.TranslateToResponseSpecificFile()
		}
	} else if errors.As(err, &StatusUtils.NotExistSource{}) {
		return response, err
	} else {
		return response, err
	}
}
