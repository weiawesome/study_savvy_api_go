package user

import (
	responseUser "study_savvy_api_go/api/response/user"
	"study_savvy_api_go/internal/repository/sql"
)

type LogoutService struct {
	Repository sql.Repository
}

func (m *LogoutService) Logout(jwt string) (responseUser.Logout, error) {
	var response responseUser.Logout
	return response, nil
}
