package mail

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	mailVerification "net/mail"
	responseUtils "study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/api/utils"
)

func validateMailVerify(mail string, code string) error {
	if mail == "" {
		return errors.New("mail can't be empty")
	} else if _, err := mailVerification.ParseAddress(mail); err != nil {
		return errors.New("mail can't parse")
	} else if code == "" {
		return errors.New("code can't be empty")
	}
	return nil
}
func MiddleWareMailVerifyContent() gin.HandlerFunc {
	return func(c *gin.Context) {
		Mail := c.Param("mail")
		Code := c.Param("code")

		err := validateMailVerify(Mail, Code)
		if err != nil {
			go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: err.Error()}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		c.Set("mail", Mail)
		c.Set("code", Code)
		c.Next()
	}
}
