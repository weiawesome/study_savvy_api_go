package nlp_edit

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"study_savvy_api_go/api/request/nlp_edit"
	"study_savvy_api_go/api/response/utils"
)

func validateNlpEditOcr(data nlp_edit.Ocr) error {
	if data.Content == "" {
		return errors.New("content can't be empty")
	}
	return nil
}
func MiddlewareNlpEditOcrContent() gin.HandlerFunc {
	return func(c *gin.Context) {

		var data nlp_edit.Ocr

		if err := c.ShouldBindJSON(&data); err != nil {
			e := utils.Error{Error: "Invalid JSON data"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}
		err := validateNlpEditOcr(data)

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
