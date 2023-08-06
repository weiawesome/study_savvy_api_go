package mail

import (
	"github.com/gin-gonic/gin"
	"net/http"
	requsetMail "study_savvy_api_go/api/request/mail"
	"study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/internal/service/mail"
)

type HandlerMailVerification struct {
	Service mail.ServiceMailVerification
}

func (h *HandlerMailVerification) Handle(c *gin.Context) {
	data, okData := c.Get("data")
	if !okData {
		e := utils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	if jsonData, ok := data.(requsetMail.Verification); ok {
		result, err := h.Service.SentVerification(jsonData)
		if err == nil {
			c.JSON(http.StatusOK, result)
		} else {
			e := utils.Error{Error: err.Error()}
			c.JSON(http.StatusInternalServerError, e)
		}
	} else {
		e := utils.Error{Error: "Internal error"}
		c.JSON(http.StatusInternalServerError, e)
	}

}
