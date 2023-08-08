package nlp_edit

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"study_savvy_api_go/api/request/nlp_edit"
	"study_savvy_api_go/api/response/utils"
)

func validateNlpEditAsr(data nlp_edit.Asr) error {
	if data.Content == "" {
		return errors.New("content can't be empty")
	}
	return nil
}
func MiddlewareNlpEditAsrContent() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Content-Type") != "application/json" {
			e := utils.Error{Error: "Content-Type must be application/json"}
			c.JSON(http.StatusUnsupportedMediaType, e)
			c.Abort()
			return
		}

		var data nlp_edit.Asr

		if err := c.ShouldBindJSON(&data); err != nil {
			e := utils.Error{Error: "Invalid JSON data"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}
		err := validateNlpEditAsr(data)

		if err != nil {
			e := utils.Error{Error: err.Error()}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}
		Id := c.Param("file_id")
		IdUuid, err := uuid.Parse(Id)

		if err != nil {
			e := utils.Error{Error: "Parameter error not uuid"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		c.Set("data", data)
		c.Set("id", IdUuid.String())
		c.Next()
	}
}
