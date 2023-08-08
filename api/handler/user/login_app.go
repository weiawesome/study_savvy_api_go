package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	userRequest "study_savvy_api_go/api/request/user"
	"study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/internal/service/user"
)

type HandlerLoginApp struct {
	Service user.ServiceLoginApp
}

func (h *HandlerLoginApp) Handle(c *gin.Context) {
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
		} else if errors.As(err, &utils.RegistrationError{}) {
			e := utils.Error{Error: err.Error()}
			c.JSON(http.StatusUnauthorized, e)
		} else {
			e := utils.Error{Error: err.Error()}
			c.JSON(http.StatusInternalServerError, e)
		}
	} else {
		e := utils.Error{Error: "Internal error"}
		c.JSON(http.StatusInternalServerError, e)
	}
}
