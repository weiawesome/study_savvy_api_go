package user

import (
	"fmt"
	"study_savvy_api_go/api/request/user"
	responseUser "study_savvy_api_go/api/response/user"
	"study_savvy_api_go/api/utils"
)

type SignupService struct{}

func (m *LoginAppService) Signup(data user.SignUp) (responseUser.Signup, error) {
	jwt, _, err := utils.GetJwt(data.Mail)
	fmt.Print(jwt)
	var result responseUser.Signup
	return result, err
}
