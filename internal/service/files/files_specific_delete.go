package files

import (
	"errors"
	"study_savvy_api_go/api/model"
	"study_savvy_api_go/api/response/files"
	"study_savvy_api_go/internal/repository/sql"
	StatusUtils "study_savvy_api_go/internal/repository/utils"
)

type ServiceFilesSpecificDelete struct {
	Repository sql.Repository
}

func (m *ServiceFilesSpecificDelete) DeleteFile(data string, id string) (files.SpecificFileDelete, error) {
	var response files.SpecificFileDelete
	File := model.File{Id: id}
	if err := m.Repository.ReadFile(&File); errors.As(err, &StatusUtils.ExistSource{}) {
		if File.UserMail != data {
			return response, StatusUtils.NotExistSource{}
		}
		return response, m.Repository.DeleteFile(File)
	} else if errors.As(err, &StatusUtils.NotExistSource{}) {
		return response, err
	} else {
		return response, err
	}
}
