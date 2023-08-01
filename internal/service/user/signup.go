package user

import (
	"gorm.io/gorm"
	"study_savvy_api_go/api/model"
	"study_savvy_api_go/api/request/user"
	responseUser "study_savvy_api_go/api/response/user"
	"study_savvy_api_go/api/utils"
)

type SignupService struct{}

func (m *LoginAppService) Signup(data user.SignUp) (responseUser.Signup, error) {
	mail := data.Mail
	var User model.User
	db := utils.GetDB()
	if result := db.First(&User, mail); result.Error == nil {

	}
	if result.Error == nil {
		// 用户已注册，邮箱已存在于数据库中
		// 可以根据需要处理逻辑，例如返回错误或做其他操作
	} else if result.Error == gorm.ErrRecordNotFound {
		// 用户未注册，邮箱不存在于数据库中
		// 可以根据需要处理逻辑，例如允许注册新用户或做其他操作
	} else {
		// 处理查询错误
	}
	var result responseUser.Signup
	return result, err
}
