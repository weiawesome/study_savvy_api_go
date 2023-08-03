package access_method

import (
	"github.com/gin-gonic/gin"
	"net/http"
	requsetAccessMethod "study_savvy_api_go/api/request/access_method"
	"study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/internal/service/access_method"
)

type HandlerAccessMethodAccessToken struct {
	Service access_method.ServiceAccessMethodAccessToken
}

func (h *HandlerAccessMethodAccessToken) Handle(c *gin.Context) {
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
		if jsonData, ok := data.(requsetAccessMethod.AccessToken); ok {
			result, err := h.Service.EditAccessToken(jsonData, stringData)
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
