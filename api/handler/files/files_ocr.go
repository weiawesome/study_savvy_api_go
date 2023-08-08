package files

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/internal/service/files"
)

type HandlerFilesOcr struct {
	Service files.ServiceFilesOcr
}

func (h *HandlerFilesOcr) Handle(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		e := utils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	page, ok := c.Get("page")
	if !ok {
		e := utils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}

	if stringDataUser, ok := user.(string); ok {
		if stringDataPage, ok := page.(int); ok {
			result, err := h.Service.GetFilesOcr(stringDataUser, stringDataPage)
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
	} else {
		e := utils.Error{Error: "Internal error"}
		c.JSON(http.StatusInternalServerError, e)
	}
}
