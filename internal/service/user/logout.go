package user

import (
	"fmt"
	responseUser "study_savvy_api_go/api/response/user"
	"study_savvy_api_go/internal/repository/redis"
)

type ServiceLogout struct {
	Repository redis.Repository
}

func (m *ServiceLogout) Logout(jwt string) (responseUser.Logout, error) {
	var response responseUser.Logout
	err := m.Repository.SetToBlacklist(jwt)
	fmt.Println(err)
	if err != nil {
		return response, err
	}
	return response, nil
}
