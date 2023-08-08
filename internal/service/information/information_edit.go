package information

import (
	"study_savvy_api_go/api/model"
	"study_savvy_api_go/api/request/information"
	responseInformation "study_savvy_api_go/api/response/information"
	"study_savvy_api_go/internal/repository/sql"
)

type ServiceInformationEdit struct {
	Repository sql.Repository
}

func (m *ServiceInformationEdit) EditInformation(data information.EditInformation, user string) (responseInformation.EditInformation, error) {
	var response responseInformation.EditInformation
	User := model.User{Mail: user, Gender: data.Gender, Name: data.Name}

	if err := m.Repository.UpdateUser(User); err == nil {
		return response, nil
	} else {
		return response, err
	}
}
