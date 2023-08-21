package mail

import (
	"github.com/gin-gonic/gin"
	"net/http"
	requsetMail "study_savvy_api_go/api/request/mail"
	responseUtils "study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/api/utils"
	"study_savvy_api_go/internal/service/mail"
)

type HandlerMailVerification struct {
	Service mail.ServiceMailVerification
}

func (h *HandlerMailVerification) Handle(c *gin.Context) {
	data, okData := c.Get("data")
	if !okData {
		go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Data not found in context"})
		e := responseUtils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	if jsonData, ok := data.(requsetMail.Verification); ok {
		result, err := h.Service.SentVerification(jsonData)
		if err == nil {
			go utils.LogInfo(utils.LogData{Event: "Success request", Method: c.Request.Method, Path: c.FullPath(), User: jsonData.Mail, Header: c.Request.Header})
			c.JSON(http.StatusOK, result)
		} else {
			go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: jsonData.Mail, Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: err.Error()}
			c.JSON(http.StatusInternalServerError, e)
		}
	} else {
		go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: jsonData.Mail, Header: c.Request.Header, Details: "Type Assertion error"})
		e := responseUtils.Error{Error: "Internal error"}
		c.JSON(http.StatusInternalServerError, e)
	}
}
