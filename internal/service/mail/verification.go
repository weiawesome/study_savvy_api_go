package mail

import (
	"errors"
	"study_savvy_api_go/api/request/mail"
	responseMial "study_savvy_api_go/api/response/mail"
	"study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/internal/repository/model"
	"study_savvy_api_go/internal/repository/redis"
	"study_savvy_api_go/internal/repository/sql"
	StatusUtils "study_savvy_api_go/internal/repository/utils"
)

type ServiceMailVerification struct {
	RedisRepository redis.Repository
	SqlRepository   sql.Repository
}

func (m *ServiceMailVerification) SentVerification(data mail.Verification) (responseMial.Verification, error) {
	var response responseMial.Verification
	User := model.User{Mail: data.Mail}

	if err := m.SqlRepository.ReadUser(&User); errors.As(err, &StatusUtils.ExistSource{}) {
		return response, utils.RegistrationError{Message: "Have been register"}
	}
	return response, m.RedisRepository.MailMission(data.Mail)
}
