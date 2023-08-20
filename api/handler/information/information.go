package information

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	responseUtils "study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/api/utils"
	"study_savvy_api_go/internal/service/information"
)

type HandlerInformation struct {
	Service information.ServiceInformation
}

func (h *HandlerInformation) Handle(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Data not found in context"})
		e := responseUtils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}

	if stringData, ok := user.(string); ok {
		result, err := h.Service.GetInformation(stringData)
		if err == nil {
			go utils.LogInfo(utils.LogData{Event: "Success request", Method: c.Request.Method, Path: c.FullPath(), User: stringData, Header: c.Request.Header})
			c.JSON(http.StatusOK, result)
		} else if errors.As(err, &responseUtils.RegistrationError{}) {
			go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringData, Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: err.Error()}
			c.JSON(http.StatusUnauthorized, e)
		} else {
			go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringData, Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: err.Error()}
			c.JSON(http.StatusInternalServerError, e)
		}
	} else {
		go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringData, Header: c.Request.Header, Details: "Type Assertion error"})
		e := responseUtils.Error{Error: "Internal error"}
		c.JSON(http.StatusInternalServerError, e)
	}
}
