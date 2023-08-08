package ai_predict

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"study_savvy_api_go/api/request/ai_predict"
	"study_savvy_api_go/api/response/utils"
)

func validateAiPredictOcrText(data ai_predict.OcrText) error {
	if data.Content == "" {
		return errors.New("content can't be empty")
	}
	return nil
}

func MiddlewareAiPredictOcrTextContent() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Content-Type") != "application/json" {
			e := utils.Error{Error: "Content-Type must be application/json"}
			c.JSON(http.StatusUnsupportedMediaType, e)
			c.Abort()
			return
		}

		var data ai_predict.OcrText

		if err := c.ShouldBindJSON(&data); err != nil {
			e := utils.Error{Error: "Invalid JSON data"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}
		err := validateAiPredictOcrText(data)

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
