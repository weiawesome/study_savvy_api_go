package user

import (
	"errors"
	"study_savvy_api_go/api/request/user"
	responseUser "study_savvy_api_go/api/response/user"
	responseUtils "study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/api/utils"
	"study_savvy_api_go/internal/repository/model"
	"study_savvy_api_go/internal/repository/sql"
	StatusUtils "study_savvy_api_go/internal/repository/utils"
)

type ServiceSignup struct {
	Repository sql.Repository
}

func (m *ServiceSignup) Signup(data user.SignUp) (responseUser.Signup, error) {
	var response responseUser.Signup

	mail := data.Mail
	password := data.Password
	gender := data.Gender
	name := data.Name

	User := model.User{Mail: mail}
	if err := m.Repository.ReadUser(&User); errors.As(err, &StatusUtils.ExistSource{}) {
		return response, responseUtils.RegistrationError{Message: "Have been registered"}
	} else if errors.As(err, &StatusUtils.NotExistSource{}) {
		salt, err := utils.GenerateSalt()
		if err != nil {
			return response, err
		}
		password = utils.GenerateHashPassword(password, salt)
		if err = m.Repository.CreateUser(model.User{Name: name, Mail: mail, Gender: gender, Password: password, Salt: salt}); err != nil {
			return response, err
		}
	} else {
		return response, err
	}

	return response, nil
}
