package access_method

import (
	"study_savvy_api_go/api/model"
	"study_savvy_api_go/api/request/access_method"
	responseAccessMethod "study_savvy_api_go/api/response/access_method"
	"study_savvy_api_go/internal/repository/sql"
)

type ServiceAccessMethodAccessToken struct {
	Repository sql.Repository
}

func (m *ServiceAccessMethodAccessToken) EditAccessToken(data access_method.AccessToken, user string) (responseAccessMethod.AccessToken, error) {
	var response responseAccessMethod.AccessToken
	AccessToken := model.AccessToken{UserMail: user}
	if err := m.Repository.FirstOrCreateAccessToken(&AccessToken); err == nil {
		AccessToken.Token = data.AccessToken
		AccessToken.AesKey = data.AesKey
		return response, m.Repository.UpdateAccessToken(AccessToken)
	} else {
		return response, err
	}
}
