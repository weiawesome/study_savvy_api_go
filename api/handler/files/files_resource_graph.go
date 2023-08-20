package files

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	responseUtils "study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/api/utils"
	StatusUtils "study_savvy_api_go/internal/repository/utils"
	"study_savvy_api_go/internal/service/files"
)

type HandlerFilesResourceGraph struct {
	Service files.ServiceFilesResourceGraph
}

func (h *HandlerFilesResourceGraph) Handle(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Data not found in context"})
		e := responseUtils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	id, ok := c.Get("id")
	if !ok {
		go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Data not found in context"})
		e := responseUtils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}

	if stringDataUser, ok := user.(string); ok {
		if stringDataId, ok := id.(string); ok {
			result, err := h.Service.GetGraph(stringDataUser, stringDataId)
			if err == nil {
				if ok := result.IsPureText(); ok {
					go utils.LogInfo(utils.LogData{Event: "Success request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataUser, Header: c.Request.Header})
					c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"error": "Pure Text"})
					return
				}
				if err := result.Exist(); err != nil {
					go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataUser, Header: c.Request.Header, Details: err.Error()})
					c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
					return
				}
				err := result.CanOpenAndSent(c)
				if err != nil {
					go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataUser, Header: c.Request.Header, Details: err.Error()})
					c.JSON(http.StatusInternalServerError, err.Error())
				}
				go utils.LogInfo(utils.LogData{Event: "Success request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataUser, Header: c.Request.Header})
			} else if errors.As(err, &StatusUtils.NotExistSource{}) {
				go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataUser, Header: c.Request.Header, Details: err.Error()})
				e := responseUtils.Error{Error: err.Error()}
				c.JSON(http.StatusNotFound, e)
			} else {
				go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: stringDataUser, Header: c.Request.Header, Details: err.Error()})
				e := responseUtils.Error{Error: err.Error()}
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
