package ai_predict

import (
	"errors"
	"github.com/google/uuid"
	"study_savvy_api_go/api/model"
	"study_savvy_api_go/api/request/ai_predict"
	responseAiPredict "study_savvy_api_go/api/response/ai_predict"
	"study_savvy_api_go/api/response/files"
	"study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/internal/repository/redis"
	"study_savvy_api_go/internal/repository/sql"
	StatusUtils "study_savvy_api_go/internal/repository/utils"
	"time"
)

type ServiceAiPredictOcrText struct {
	SqlRepository   sql.Repository
	RedisRepository redis.Repository
}

func (m *ServiceAiPredictOcrText) ExecuteOcrText(data ai_predict.OcrText, user string) (responseAiPredict.OcrText, error) {
	var response responseAiPredict.OcrText
	User := model.User{Mail: user}
	Preloads := []string{model.ApiKeyRelate(), model.AccessTokenRelate()}

	if err := m.SqlRepository.PreLoadReadUser(&User, Preloads); errors.As(err, &StatusUtils.ExistSource{}) {
		Id := uuid.New().String()
		File := model.File{Id: Id, UserMail: user, Status: "PENDING", CreatedAt: time.Now(), Resource: files.PureText, Type: "OCR"}
		if err := m.SqlRepository.CreateFile(File); err == nil {
			return response, m.RedisRepository.OcrTextMission(Id, data.Content, data.Prompt, User.ApiKey, User.AccessToken)
		} else {
			return response, err
		}
	} else if errors.As(err, &StatusUtils.ExistSource{}) {
		return response, utils.RegistrationError{Message: "Have not register"}
	} else {
		return response, err
	}
}
