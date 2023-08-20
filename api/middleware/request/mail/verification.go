package mail

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	mailVerification "net/mail"
	"study_savvy_api_go/api/request/mail"
	responseUtils "study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/api/utils"
)

func validateMailVerification(data mail.Verification) error {
	if data.Mail == "" {
		return errors.New("mail can't be empty")
	} else if _, err := mailVerification.ParseAddress(data.Mail); err != nil {
		return errors.New("mail can't parse")
	}
	return nil
}
func MiddlewareMailVerificationContent() gin.HandlerFunc {
	return func(c *gin.Context) {

		var data mail.Verification

		if err := c.ShouldBindJSON(&data); err != nil {
			go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: "Invalid JSON data"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}
		err := validateMailVerification(data)

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
