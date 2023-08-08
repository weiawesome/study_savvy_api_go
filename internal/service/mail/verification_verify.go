package mail

import (
	responseMial "study_savvy_api_go/api/response/mail"
	"study_savvy_api_go/internal/repository/redis"
)

type ServiceMailVerify struct {
	RedisRepository redis.Repository
}

func (m *ServiceMailVerify) Verify(user string, code string) (responseMial.VerifyMail, error) {
	var response responseMial.VerifyMail
	return response, m.RedisRepository.ValidateInVerification(user, code)
}
