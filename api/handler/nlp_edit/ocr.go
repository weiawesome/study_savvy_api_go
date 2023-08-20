package nlp_edit

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	requsetNlpEdit "study_savvy_api_go/api/request/nlp_edit"
	responseUtils "study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/api/utils"
	"study_savvy_api_go/internal/service/nlp_edit"
)

type HandlerNlpEditOcr struct {
	Service nlp_edit.ServiceNlpEditOcr
}

func (h *HandlerNlpEditOcr) Handle(c *gin.Context) {
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
	id, okId := c.Get("id")
	if !okId {
		go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Data not found in context"})
		e := responseUtils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	if stringDataUser, ok := user.(string); ok {
		if stringDataId, ok := id.(string); ok {
			if jsonData, ok := data.(requsetNlpEdit.Ocr); ok {
				result, err := h.Service.ExecuteOcr(jsonData, stringDataUser, stringDataId)
				if err == nil {
					go utils.LogInfo(utils.LogData{Event: "Success request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataUser, Header: c.Request.Header, Content: jsonData})
					c.JSON(http.StatusOK, result)
				} else if errors.As(err, &responseUtils.RegistrationError{}) {
					go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataUser, Header: c.Request.Header, Content: jsonData, Details: err.Error()})
					e := responseUtils.Error{Error: err.Error()}
					c.JSON(http.StatusUnauthorized, e)
				} else if errors.As(err, &responseUtils.ExistError{}) {
					go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataUser, Header: c.Request.Header, Content: jsonData, Details: err.Error()})
					e := responseUtils.Error{Error: err.Error()}
					c.JSON(http.StatusNotFound, e)
				} else if errors.As(err, &responseUtils.AuthError{}) {
					go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataUser, Header: c.Request.Header, Content: jsonData, Details: err.Error()})
					e := responseUtils.Error{Error: err.Error()}
					c.JSON(http.StatusUnprocessableEntity, e)
				} else {
					go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataUser, Header: c.Request.Header, Content: jsonData, Details: err.Error()})
					e := responseUtils.Error{Error: err.Error()}
					c.JSON(http.StatusInternalServerError, e)
				}
			} else {
				go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataUser, Header: c.Request.Header, Content: jsonData, Details: "Type Assertion error"})
				e := responseUtils.Error{Error: "Internal error"}
				c.JSON(http.StatusInternalServerError, e)
			}
		} else {
			go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataUser, Header: c.Request.Header, Details: "Type Assertion error"})
			e := responseUtils.Error{Error: "Internal error"}
			c.JSON(http.StatusInternalServerError, e)
		}
	} else {
		go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataUser, Header: c.Request.Header, Details: "Type Assertion error"})
		e := responseUtils.Error{Error: "Internal error"}
		c.JSON(http.StatusInternalServerError, e)
	}
}
