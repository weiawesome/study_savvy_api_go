package information

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"study_savvy_api_go/api/request/information"
	"study_savvy_api_go/api/response/utils"
)

func validatePasswordEdit(data information.EditPassword) error {
	if data.CurrentPassword == "" || data.EditPassword == "" {
		return errors.New("password can't be empty")
	} else if len(data.CurrentPassword) < 8 || len(data.EditPassword) < 8 {
		return errors.New("length of password can't shorten than 8 chars")
	} else if len(data.CurrentPassword) > 30 || len(data.EditPassword) > 30 {
		return errors.New("length of password can't larger than 30 chars")
	} else if strings.ContainsAny(data.CurrentPassword, " ") || strings.ContainsAny(data.EditPassword, " ") {
		return errors.New("password can't contain space")
	}
	return nil
}

func MiddleWarePasswordEditContent() gin.HandlerFunc {
	return func(c *gin.Context) {

		var data information.EditPassword

		if err := c.ShouldBindJSON(&data); err != nil {
			e := utils.Error{Error: "Invalid JSON data"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}
		err := validatePasswordEdit(data)

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
