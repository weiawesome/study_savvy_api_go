package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/mail"
	"strings"
	"study_savvy_api_go/api/request/user"
	"study_savvy_api_go/api/response/utils"
)

func validateSignup(data user.SignUp) error {
	if data.Mail == "" {
		return errors.New("mail can't be empty")
	} else if data.Password == "" {
		return errors.New("password can't be empty")
	} else if len(data.Password) < 8 {
		return errors.New("length of password can't shorten than 8 chars")
	} else if strings.ContainsAny(data.Password, " ") {
		return errors.New("password can't contain space")
	} else if _, err := mail.ParseAddress(data.Mail); err != nil {
		return errors.New("mail can't parse")
	} else if !(data.Gender == "male" || data.Gender == "female" || data.Gender == "other") {
		return errors.New("gender error type")
	}
	return nil
}

func MiddleWareSignupContent() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Content-Type") != "application/json" {
			e := utils.Error{Error: "Content-Type must be application/json"}
			c.JSON(http.StatusUnsupportedMediaType, e)
			c.Abort()
			return
		}

		var data user.SignUp
		if err := c.ShouldBindJSON(&data); err != nil {
			e := utils.Error{Error: "Invalid JSON data"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}
		err := validateSignup(data)

		if err != nil {
			e := utils.Error{Error: err.Error()}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}
		c.Set("data", data)
		c.Next()
	}
}
