package files

import (
	"errors"
	"study_savvy_api_go/api/model"
	"study_savvy_api_go/api/response/files"
	"study_savvy_api_go/internal/repository/sql"
	StatusUtils "study_savvy_api_go/internal/repository/utils"
)

type ServiceFilesResourceAudio struct {
	Repository sql.Repository
}

func (m *ServiceFilesResourceAudio) GetAudio(data string, id string) (files.AudioFile, error) {
	var response files.AudioFile
	File := model.File{UserMail: data, Id: id}
	if err := m.Repository.ReadFile(&File); errors.As(err, &StatusUtils.ExistSource{}) {
		if File.UserMail != data {
			return response, StatusUtils.NotExistSource{}
		}
		return files.AudioFile{FilePath: File.Resource}, nil
	} else if errors.As(err, &StatusUtils.NotExistSource{}) {
		return response, err
	} else {
		return response, err
	}
}
