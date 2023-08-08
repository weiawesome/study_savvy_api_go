package user

import (
	"errors"
	"study_savvy_api_go/api/model"
	"study_savvy_api_go/api/request/user"
	responseUser "study_savvy_api_go/api/response/user"
	responseUtils "study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/api/utils"
	"study_savvy_api_go/internal/repository/sql"
	StatusUtils "study_savvy_api_go/internal/repository/utils"
)

type ServiceLoginWeb struct {
	Repository sql.Repository
}

func (m *ServiceLoginWeb) Login(data user.LoginWeb) (responseUser.LoginWeb, responseUser.LoginToken, error) {
	var response responseUser.LoginWeb
	var responseToken responseUser.LoginToken
	User := model.User{Mail: data.Mail}

	if err := m.Repository.ReadUser(&User); errors.As(err, &StatusUtils.ExistSource{}) {
		if utils.ValidatePassword(data.Password, User.Password, User.Salt) {
			jwt, csrf, err := utils.GetJwt(data.Mail)
			return responseUser.LoginWeb{}, responseUser.LoginToken{JwtToken: jwt, CsrfToken: csrf}, err
		} else {
			return response, responseToken, responseUtils.RegistrationError{Message: "Password error"}
		}
	} else if errors.As(err, &StatusUtils.NotExistSource{}) {
		return response, responseToken, responseUtils.RegistrationError{Message: "Have not sign up"}
	} else {
		return response, responseToken, err
	}
}
