package access_method

import (
	"study_savvy_api_go/api/model"
	"study_savvy_api_go/api/request/access_method"
	responseAccessMethod "study_savvy_api_go/api/response/access_method"
	"study_savvy_api_go/internal/repository/sql"
)

type ServiceAccessMethodApiKey struct {
	Repository sql.Repository
}

func (m *ServiceAccessMethodApiKey) EditApiKey(data access_method.ApiKey, user string) (responseAccessMethod.ApiKey, error) {
	var response responseAccessMethod.ApiKey
	ApiKey := model.ApiKey{UserMail: user}
	if err := m.Repository.FirstOrCreateApiKey(&ApiKey); err == nil {
		ApiKey.Key = data.ApiKey
		ApiKey.AesKey = data.AesKey
		return response, m.Repository.UpdateApiKey(ApiKey)
	} else {
		return response, err
	}
}
