package access_method

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"study_savvy_api_go/api/request/access_method"
	"study_savvy_api_go/api/response/utils"
)

func validateAccessTokenEdit(data access_method.AccessToken) error {
	if data.AccessToken == "" {
		return errors.New("AccessToken can't be empty")
	} else if data.AesKey == "" {
		return errors.New("AesKey can't be empty")
	}
	return nil
}

func MiddleWareAccessTokenEditContent() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Content-Type") != "application/json" {
			e := utils.Error{Error: "Content-Type must be application/json"}
			c.JSON(http.StatusUnsupportedMediaType, e)
			c.Abort()
			return
		}

		var data access_method.AccessToken

		if err := c.ShouldBindJSON(&data); err != nil {
			e := utils.Error{Error: "Invalid JSON data"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}
		err := validateAccessTokenEdit(data)

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
