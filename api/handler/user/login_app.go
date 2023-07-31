package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	userRequest "study_savvy_api_go/api/request/user"
	"study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/internal/service/user"
)

type LoginAppHandler struct {
	Service user.LoginAppService
}

func (h *LoginAppHandler) Handle(c *gin.Context) {
	data, ok := c.Get("data")
	if !ok {
		e := utils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}

	if jsonData, ok := data.(userRequest.LoginApp); ok {
		result, err := h.Service.Login(jsonData)
		if err == nil {
			c.JSON(http.StatusOK, result)
		} else {
			e := utils.Error{Error: "Data type mismatch"}
			c.JSON(http.StatusInternalServerError, e)
		}
	} else {
		e := utils.Error{Error: "Internal error"}
		c.JSON(http.StatusInternalServerError, e)
	}
}
