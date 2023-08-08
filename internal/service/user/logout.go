package user

import (
	responseUser "study_savvy_api_go/api/response/user"
	"study_savvy_api_go/internal/repository/redis"
)

type ServiceLogout struct {
	Repository redis.Repository
}

func (m *ServiceLogout) Logout(jti string) (responseUser.Logout, error) {
	var response responseUser.Logout
	err := m.Repository.SetToBlacklist(jti)

	return response, err
}
