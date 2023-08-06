package mail

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/internal/service/mail"
)

type HandlerMailVerify struct {
	Service mail.ServiceMailVerify
}

func (h *HandlerMailVerify) Handle(c *gin.Context) {
	user, okUser := c.Get("mail")
	if !okUser {
		e := utils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	code, okCode := c.Get("code")
	if !okCode {
		e := utils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	if userData, ok := user.(string); ok {
		if codeData, ok := code.(string); ok {
			result, err := h.Service.Verify(userData, codeData)
			if err == nil {
				c.JSON(http.StatusOK, result)
			} else {
				e := utils.Error{Error: err.Error()}
				c.JSON(http.StatusBadRequest, e)
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
