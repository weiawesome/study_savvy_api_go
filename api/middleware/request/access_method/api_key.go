package access_method

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"study_savvy_api_go/api/request/access_method"
	"study_savvy_api_go/api/response/utils"
)

func validateApiKeyEdit(data access_method.ApiKey) error {
	if data.ApiKey == "" {
		return errors.New("ApiKey can't be empty")
	} else if data.AesKey == "" {
		return errors.New("AesKey can't be empty")
	}
	return nil
}

func MiddleWareApiKeyEditContent() gin.HandlerFunc {
	return func(c *gin.Context) {

		var data access_method.ApiKey

		if err := c.ShouldBindJSON(&data); err != nil {
			e := utils.Error{Error: "Invalid JSON data"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}
		err := validateApiKeyEdit(data)

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
