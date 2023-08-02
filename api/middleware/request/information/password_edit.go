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
	if data.OriginalPwd == "" || data.NewPwd == "" {
		return errors.New("password can't be empty")
	} else if len(data.OriginalPwd) < 8 || len(data.NewPwd) < 8 {
		return errors.New("length of password can't shorten than 8 chars")
	} else if strings.ContainsAny(data.OriginalPwd, " ") || strings.ContainsAny(data.NewPwd, " ") {
		return errors.New("password can't contain space")
	}
	return nil
}

func MiddleWarePasswordEditContent() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Content-Type") != "application/json" {
			e := utils.Error{Error: "Content-Type must be application/json"}
			c.JSON(http.StatusUnsupportedMediaType, e)
			c.Abort()
			return
		}

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
