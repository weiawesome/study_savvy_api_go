package nlp_edit

import (
	"errors"
	"study_savvy_api_go/api/request/nlp_edit"
	responseNlpEdit "study_savvy_api_go/api/response/nlp_edit"
	"study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/internal/repository/model"
	"study_savvy_api_go/internal/repository/redis"
	"study_savvy_api_go/internal/repository/sql"
	StatusUtils "study_savvy_api_go/internal/repository/utils"
)

type ServiceNlpEditAsr struct {
	SqlRepository   sql.Repository
	RedisRepository redis.Repository
}

func (m *ServiceNlpEditAsr) ExecuteAsr(data nlp_edit.Asr, user string, id string) (responseNlpEdit.Asr, error) {
	var response responseNlpEdit.Asr
	User := model.User{Mail: user}
	File := model.File{Id: id}
	Preloads := []string{model.ApiKeyRelate(), model.AccessTokenRelate()}
	if err := m.SqlRepository.ReadFile(&File); errors.As(err, &StatusUtils.ExistSource{}) {
		if File.UserMail == User.Mail {
			if err := m.SqlRepository.PreLoadReadUser(&User, Preloads); errors.As(err, &StatusUtils.ExistSource{}) {
				File.Status = "PENDING"
				if err := m.SqlRepository.UpdateFile(File); err == nil {
					return response, m.RedisRepository.NlpEditAsrMission(File.Id, data.Content, data.Prompt, User.ApiKey, User.AccessToken)
				} else {
					return response, err
				}
			} else if errors.As(err, &StatusUtils.NotExistSource{}) {
				return response, utils.RegistrationError{Message: "Have not register"}
			} else {
				return response, err
			}
		} else {
			return response, utils.AuthError{Message: "Auth error"}
		}
	} else if errors.As(err, &StatusUtils.NotExistSource{}) {
		return response, utils.ExistError{Message: "No such file with id"}
	} else {
		return response, err
	}
}
