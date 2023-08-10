package information

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"study_savvy_api_go/api/request/information"
	"study_savvy_api_go/api/response/utils"
)

func validateInformationEdit(data information.EditInformation) error {
	if data.Name == "" {
		return errors.New("name can't be empty")
	} else if len(data.Name) > 30 {
		return errors.New("length of name can't large than 30 chars")
	} else if !(data.Gender == "male" || data.Gender == "female" || data.Gender == "other") {
		return errors.New("gender error type")
	}
	return nil
}

func MiddleWareInformationEditContent() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Content-Type") != "application/json" {
			e := utils.Error{Error: "Content-Type must be application/json"}
			c.JSON(http.StatusUnsupportedMediaType, e)
			c.Abort()
			return
		}

		var data information.EditInformation

		if err := c.ShouldBindJSON(&data); err != nil {
			e := utils.Error{Error: "Invalid JSON data"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}
		err := validateInformationEdit(data)

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
