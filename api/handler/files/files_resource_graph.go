package files

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"study_savvy_api_go/api/response/utils"
	StatusUtils "study_savvy_api_go/internal/repository/utils"
	"study_savvy_api_go/internal/service/files"
)

type HandlerFilesResourceGraph struct {
	Service files.ServiceFilesResourceGraph
}

func (h *HandlerFilesResourceGraph) Handle(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		e := utils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	id, ok := c.Get("id")
	if !ok {
		e := utils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}

	if stringDataUser, ok := user.(string); ok {
		if stringDataId, ok := id.(string); ok {
			result, err := h.Service.GetGraph(stringDataUser, stringDataId)
			if err == nil {
				if err := result.Exist(); err != nil {
					c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
					return
				}
				err := result.CanOpenAndSent(c)
				if err != nil {
					c.JSON(http.StatusInternalServerError, err.Error())
				}
			} else if errors.As(err, &StatusUtils.NotExistSource{}) {
				e := utils.Error{Error: err.Error()}
				c.JSON(http.StatusNotFound, e)
			} else {
				e := utils.Error{Error: err.Error()}
				c.JSON(http.StatusInternalServerError, e)
			}
		}
	} else {
		e := utils.Error{Error: "Internal error"}
		c.JSON(http.StatusInternalServerError, e)
	}
}
