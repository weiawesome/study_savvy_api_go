package ai_predict

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	requsetAiPredict "study_savvy_api_go/api/request/ai_predict"
	"study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/internal/service/ai_predict"
)

type HandlerAiPredictOcr struct {
	Service ai_predict.ServiceAiPredictOcr
}

func (h *HandlerAiPredictOcr) Handle(c *gin.Context) {
	user, okUser := c.Get("user")
	if !okUser {
		e := utils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	data, okData := c.Get("data")
	if !okData {
		e := utils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	if stringData, ok := user.(string); ok {
		if jsonData, ok := data.(requsetAiPredict.Ocr); ok {
			result, err := h.Service.ExecuteOcr(jsonData, stringData)
			if err == nil {
				c.JSON(http.StatusOK, result)
			} else if errors.As(err, &utils.RegistrationError{}) {
				e := utils.Error{Error: err.Error()}
				c.JSON(http.StatusUnauthorized, e)
			} else {
				e := utils.Error{Error: err.Error()}
				c.JSON(http.StatusInternalServerError, e)
			}
		} else {
			e := utils.Error{Error: "Internal error"}
			c.JSON(http.StatusInternalServerError, e)
		}
	} else {
		e := utils.Error{Error: "Internal error"}
		c.JSON(http.StatusInternalServerError, e)
	}
}
