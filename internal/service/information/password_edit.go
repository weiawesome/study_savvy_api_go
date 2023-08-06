package information

import (
	"errors"
	"study_savvy_api_go/api/model"
	"study_savvy_api_go/api/request/information"
	responseInformation "study_savvy_api_go/api/response/information"
	responseUtils "study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/api/utils"
	"study_savvy_api_go/internal/repository/sql"
	StatusUtils "study_savvy_api_go/internal/repository/utils"
)

type ServicePasswordEdit struct {
	Repository sql.Repository
}

func (m *ServicePasswordEdit) EditPassword(data information.EditPassword, user string) (responseInformation.EditPassword, error) {
	var response responseInformation.EditPassword

	User := model.User{Mail: user}

	if err := m.Repository.ReadUser(&User); errors.As(err, &StatusUtils.ExistSource{}) {
		if utils.ValidatePassword(data.OriginalPwd, User.Password, User.Salt) {
			salt, err := utils.GenerateSalt()
			if err != nil {
				return response, err
			}
			password := utils.GenerateHashPassword(data.NewPwd, salt)
			return response, m.Repository.UpdateUser(model.User{Mail: user, Password: password, Salt: salt})
		} else {
			return response, responseUtils.RegistrationError{Message: "Password error"}
		}
	} else if errors.As(err, &StatusUtils.NotExistSource{}) {
		return response, responseUtils.RegistrationError{Message: "Have not sign up"}
	} else {
		return response, err
	}
}
