package nlp_edit

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	requsetNlpEdit "study_savvy_api_go/api/request/nlp_edit"
	responseUtils "study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/api/utils"
	"study_savvy_api_go/internal/service/logger"
	"study_savvy_api_go/internal/service/nlp_edit"
)

type HandlerNlpEditAsr struct {
	Service    nlp_edit.ServiceNlpEditAsr
	LogService logger.ServiceLogger
}

func (h *HandlerNlpEditAsr) Handle(c *gin.Context) {
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
			if jsonData, ok := data.(requsetNlpEdit.Asr); ok {
				result, err := h.Service.ExecuteAsr(jsonData, stringDataUser, stringDataId)
				if err == nil {
					go h.LogService.Info(utils.LogData{Event: "Success request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataUser, Header: c.Request.Header})
					c.JSON(http.StatusOK, result)
				} else if errors.As(err, &responseUtils.RegistrationError{}) {
					go h.LogService.Warn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataUser, Header: c.Request.Header, Details: err.Error()})
					e := responseUtils.Error{Error: err.Error()}
					c.JSON(http.StatusUnauthorized, e)
				} else if errors.As(err, &responseUtils.ExistError{}) {
					go h.LogService.Warn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataUser, Header: c.Request.Header, Details: err.Error()})
					e := responseUtils.Error{Error: err.Error()}
					c.JSON(http.StatusNotFound, e)
				} else if errors.As(err, &responseUtils.AuthError{}) {
					go h.LogService.Warn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataUser, Header: c.Request.Header, Details: err.Error()})
					e := responseUtils.Error{Error: err.Error()}
					c.JSON(http.StatusUnprocessableEntity, e)
				} else {
					go h.LogService.Error(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataUser, Header: c.Request.Header, Details: err.Error()})
					e := responseUtils.Error{Error: err.Error()}
					c.JSON(http.StatusInternalServerError, e)
				}
			} else {
				go h.LogService.Error(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataUser, Header: c.Request.Header, Details: "Type Assertion error"})
				e := responseUtils.Error{Error: "Internal error"}
				c.JSON(http.StatusInternalServerError, e)
			}
		} else {
			go h.LogService.Error(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataUser, Header: c.Request.Header, Details: "Type Assertion error"})
			e := responseUtils.Error{Error: "Internal error"}
			c.JSON(http.StatusInternalServerError, e)
		}
	} else {
		go h.LogService.Error(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataUser, Header: c.Request.Header, Details: "Type Assertion error"})
		e := responseUtils.Error{Error: "Internal error"}
		c.JSON(http.StatusInternalServerError, e)
	}
}
