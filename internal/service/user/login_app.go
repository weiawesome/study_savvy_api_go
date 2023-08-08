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

type ServiceLoginApp struct {
	Repository sql.Repository
}

func (m *ServiceLoginApp) Login(data user.LoginApp) (responseUser.LoginApp, error) {
	var response responseUser.LoginApp
	User := model.User{Mail: data.Mail}

	if err := m.Repository.ReadUser(&User); errors.As(err, &StatusUtils.ExistSource{}) {
		if utils.ValidatePassword(data.Password, User.Password, User.Salt) {
			jwt, _, err := utils.GetJwt(data.Mail)
			return responseUser.LoginApp{Token: jwt}, err
		} else {
			return response, responseUtils.RegistrationError{Message: "Password error"}
		}
	} else if errors.As(err, &StatusUtils.NotExistSource{}) {
		return response, responseUtils.RegistrationError{Message: "Have not sign up"}
	} else {
		return response, err
	}
}
