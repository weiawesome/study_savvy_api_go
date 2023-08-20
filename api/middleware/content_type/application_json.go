package content_type

import (
	"github.com/gin-gonic/gin"
	"net/http"
	responseUtils "study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/api/utils"
)

func MiddleWareApplicationJson() gin.HandlerFunc {
	return func(c *gin.Context) {

		contentType := c.ContentType()

		if contentType != "application/json" {
			go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Content-Type not application/json"})
			e := responseUtils.Error{Error: "Content-Type must be application/json not " + contentType}
			c.JSON(http.StatusUnsupportedMediaType, e)
			c.Abort()
			return
		}

		c.Next()
	}
}
