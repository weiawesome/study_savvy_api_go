package mail

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	mailVerification "net/mail"
	"study_savvy_api_go/api/request/mail"
	"study_savvy_api_go/api/response/utils"
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
			e := utils.Error{Error: "Invalid JSON data"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}
		err := validateMailVerification(data)

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
