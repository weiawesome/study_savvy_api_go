package information

import (
	"github.com/gin-gonic/gin"
	"net/http"
	requsetInformation "study_savvy_api_go/api/request/information"
	"study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/internal/service/information"
)

type HandlerInformationEdit struct {
	Service information.ServiceInformationEdit
}

func (h *HandlerInformationEdit) Handle(c *gin.Context) {
	user, okUser := c.Get("user")
	if !okUser {
		e := utils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	data, okData := c.Get("data")
	if !okData {
		e := utils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}

	if stringData, ok := user.(string); ok {
		if jsonData, ok := data.(requsetInformation.EditInformation); ok {
			result, err := h.Service.EditInformation(jsonData, stringData)
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
