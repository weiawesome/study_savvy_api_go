package user

import (
	"study_savvy_api_go/api/request/user"
	responseUser "study_savvy_api_go/api/response/user"
	"study_savvy_api_go/api/utils"
)

type LoginAppService struct{}

func (m *LoginAppService) Login(data user.LoginApp) (responseUser.LoginApp, error) {
	jwt, _, err := utils.GetJwt(data.Mail)
	return responseUser.LoginApp{Token: jwt}, err
}
