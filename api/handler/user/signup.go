package user

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	userRequest "study_savvy_api_go/api/request/user"
	"study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/internal/service/user"
)

type SignupHandler struct {
	Service user.SignupService
}

func (h *SignupHandler) Handle(c *gin.Context) {
	data, ok := c.Get("data")
	if !ok {
		e := utils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}

	if jsonData, ok := data.(userRequest.SignUp); ok {
		result, err := h.Service.Signup(jsonData)
		if err == nil {
			c.JSON(http.StatusOK, result)
		} else if errors.As(err, &utils.RegistrationError{}) {
			e := utils.Error{Error: err.Error()}
			fmt.Print("Have Been signup")
			c.JSON(http.StatusUnauthorized, e)
		} else {
			fmt.Println(err)
			e := utils.Error{Error: "Data type mismatch"}
			c.JSON(http.StatusInternalServerError, e)
		}
	} else {
		e := utils.Error{Error: "Internal error"}
		c.JSON(http.StatusInternalServerError, e)
	}
}