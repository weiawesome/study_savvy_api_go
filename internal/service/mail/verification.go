package mail

import (
	"study_savvy_api_go/api/request/mail"
	responseMial "study_savvy_api_go/api/response/mail"
	"study_savvy_api_go/internal/repository/redis"
)

type ServiceMailVerification struct {
	RedisRepository redis.Repository
}

func (m *ServiceMailVerification) SentVerification(data mail.Verification) (responseMial.Verification, error) {
	var response responseMial.Verification
	return response, m.RedisRepository.MailMission(data.Mail)
}
