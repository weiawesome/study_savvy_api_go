package files

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"study_savvy_api_go/api/response/utils"
	StatusUtils "study_savvy_api_go/internal/repository/utils"
	"study_savvy_api_go/internal/service/files"
)

type HandlerFilesSpecific struct {
	Service files.ServiceFilesSpecific
}

func (h *HandlerFilesSpecific) Handle(c *gin.Context) {
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

	if stringData, ok := user.(string); ok {
		if stringDataId, ok := id.(string); ok {
			result, err := h.Service.GetFile(stringData, stringDataId)
			if err == nil {
				c.JSON(http.StatusOK, result)
			} else if errors.As(err, &StatusUtils.NotExistSource{}) {
				e := utils.Error{Error: err.Error()}
				c.JSON(http.StatusNotFound, e)
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
