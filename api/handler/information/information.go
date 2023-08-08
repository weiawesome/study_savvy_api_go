package information

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/internal/service/information"
)

type HandlerInformation struct {
	Service information.ServiceInformation
}

func (h *HandlerInformation) Handle(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		e := utils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}

	if stringData, ok := user.(string); ok {
		result, err := h.Service.GetInformation(stringData)
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
