package nlp_edit

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	requsetNlpEdit "study_savvy_api_go/api/request/nlp_edit"
	"study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/internal/service/nlp_edit"
)

type HandlerNlpEditOcr struct {
	Service nlp_edit.ServiceNlpEditOcr
}

func (h *HandlerNlpEditOcr) Handle(c *gin.Context) {
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
	id, okId := c.Get("id")
	if !okId {
		e := utils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	if stringDataUser, ok := user.(string); ok {
		if stringDataId, ok := id.(string); ok {
			if jsonData, ok := data.(requsetNlpEdit.Ocr); ok {
				result, err := h.Service.ExecuteOcr(jsonData, stringDataUser, stringDataId)
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
		} else {
			e := utils.Error{Error: "Internal error"}
			c.JSON(http.StatusInternalServerError, e)
		}
	} else {
		e := utils.Error{Error: "Internal error"}
		c.JSON(http.StatusInternalServerError, e)
	}
}
