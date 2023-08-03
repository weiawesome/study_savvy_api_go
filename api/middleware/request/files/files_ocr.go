package files

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"study_savvy_api_go/api/response/utils"
)

func MiddleWareFilesOcrContent() gin.HandlerFunc {
	return func(c *gin.Context) {
		page := c.DefaultQuery("page", "1")
		pageInt, err := strconv.Atoi(page)

		if err != nil {
			e := utils.Error{Error: "Parameter error"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		c.Set("page", pageInt)
		c.Next()
	}
}
