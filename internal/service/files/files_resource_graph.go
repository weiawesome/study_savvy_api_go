package files

import (
	"errors"
	"study_savvy_api_go/api/model"
	"study_savvy_api_go/api/response/files"
	"study_savvy_api_go/internal/repository/sql"
	StatusUtils "study_savvy_api_go/internal/repository/utils"
)

type ServiceFilesResourceGraph struct {
	Repository sql.Repository
}

func (m *ServiceFilesResourceGraph) GetGraph(data string, id string) (files.GraphFile, error) {
	var response files.GraphFile
	File := model.File{UserMail: data, Id: id}
	if err := m.Repository.ReadFile(&File); errors.As(err, &StatusUtils.ExistSource{}) {
		return files.GraphFile{FilePath: File.Source}, nil
	} else if errors.As(err, &StatusUtils.NotExistSource{}) {
		return response, err
	} else {
		return response, err
	}
}
