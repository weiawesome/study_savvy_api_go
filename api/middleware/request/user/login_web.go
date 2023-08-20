package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/mail"
	"strings"
	"study_savvy_api_go/api/request/user"
	responseUtils "study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/api/utils"
)

func validateLoginWeb(data user.LoginWeb) error {
	if data.Mail == "" {
		return errors.New("mail can't be empty")
	} else if data.Password == "" {
		return errors.New("password can't be empty")
	} else if len(data.Password) < 8 {
		return errors.New("length of password can't shorten than 8 chars")
	} else if len(data.Password) > 30 {
		return errors.New("length of password can't larger than 30 chars")
	} else if strings.ContainsAny(data.Password, " ") {
		return errors.New("password can't contain space")
	} else if _, err := mail.ParseAddress(data.Mail); err != nil {
		return errors.New("mail can't parse")
	}
	return nil
}

func MiddleWareLoginWebContent() gin.HandlerFunc {
	return func(c *gin.Context) {

		var data user.LoginWeb

		if err := c.ShouldBindJSON(&data); err != nil {
			go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: "Invalid JSON data"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}
		err := validateLoginWeb(data)

		if err != nil {
			go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: err.Error()}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		c.Set("data", data)
		c.Next()
	}
}
