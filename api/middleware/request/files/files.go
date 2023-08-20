package files

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	responseUtils "study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/api/utils"
)

func MiddleWareFilesContent() gin.HandlerFunc {
	return func(c *gin.Context) {
		page := c.DefaultQuery("page", "1")
		pageInt, err := strconv.Atoi(page)

		if err != nil {
			go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: "Parameter error"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		c.Set("page", pageInt)
		c.Next()
	}
}
