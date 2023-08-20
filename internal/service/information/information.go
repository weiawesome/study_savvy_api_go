package information

import (
	"errors"
	"study_savvy_api_go/api/response/information"
	responseUtils "study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/internal/repository/model"
	"study_savvy_api_go/internal/repository/sql"
	StatusUtils "study_savvy_api_go/internal/repository/utils"
)

type ServiceInformation struct {
	Repository sql.Repository
}

func (m *ServiceInformation) GetInformation(data string) (information.Information, error) {
	var response information.Information
	User := model.User{Mail: data}

	if err := m.Repository.ReadUser(&User); errors.As(err, &StatusUtils.ExistSource{}) {
		return information.Information{Name: User.Name, Gender: User.Gender, Mail: User.Mail}, nil
	} else if errors.As(err, &StatusUtils.NotExistSource{}) {
		return response, responseUtils.RegistrationError{Message: "Have not sign up"}
	} else {
		return response, err
	}
}
