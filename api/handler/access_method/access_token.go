package access_method

import (
	"github.com/gin-gonic/gin"
	"net/http"
	requsetAccessMethod "study_savvy_api_go/api/request/access_method"
	responseUtils "study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/api/utils"
	"study_savvy_api_go/internal/service/access_method"
)

type HandlerAccessMethodAccessToken struct {
	Service access_method.ServiceAccessMethodAccessToken
}

func (h *HandlerAccessMethodAccessToken) Handle(c *gin.Context) {
	user, okUser := c.Get("user")
	if !okUser {
		go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Data not found in context"})
		e := responseUtils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	data, okData := c.Get("data")
	if !okData {
		go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Data not found in context"})
		e := responseUtils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	if stringData, ok := user.(string); ok {
		if jsonData, ok := data.(requsetAccessMethod.AccessToken); ok {
			result, err := h.Service.EditAccessToken(jsonData, stringData)
			if err == nil {
				go utils.LogInfo(utils.LogData{Event: "Success request", Method: c.Request.Method, Path: c.FullPath(), User: stringData, Header: c.Request.Header, Content: jsonData})
				c.JSON(http.StatusOK, result)
			} else {
				go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringData, Header: c.Request.Header, Content: jsonData, Details: err.Error()})
				e := responseUtils.Error{Error: err.Error()}
				c.JSON(http.StatusInternalServerError, e)
			}
		} else {
			go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringData, Header: c.Request.Header, Content: jsonData, Details: "Type Assertion error"})
			e := responseUtils.Error{Error: "Internal error"}
			c.JSON(http.StatusInternalServerError, e)
		}
	} else {
		go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringData, Header: c.Request.Header, Details: "Type Assertion error"})
		e := responseUtils.Error{Error: "Internal error"}
		c.JSON(http.StatusInternalServerError, e)
	}
}
