package user

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"study_savvy_api_go/api/model"
	"study_savvy_api_go/api/request/user"
	responseUser "study_savvy_api_go/api/response/user"
	responseUtils "study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/api/utils"
)

type SignupService struct{}

func (m *SignupService) Signup(data user.SignUp) (responseUser.Signup, error) {
	var response responseUser.Signup

	mail := data.Mail
	password := data.Password
	gender := data.Gender
	name := data.Name

	User := model.User{Mail: mail}
	db := utils.GetDB()
	fmt.Println(mail)
	if result := db.First(&User); result.Error == nil {
		return response, responseUtils.RegistrationError{Message: "Have been registered"}
	} else if errors.As(result.Error, &gorm.ErrRecordNotFound) {
		salt, err := utils.GenerateSalt()
		if err != nil {
			return response, err
		}
		password = utils.GenerateHashPassword(password, salt)
		db.Create(&model.User{Name: name, Mail: mail, Gender: gender, Password: password, Salt: salt})
	} else {
		fmt.Println(result.Error)
		return response, result.Error
	}

	return response, nil
}
