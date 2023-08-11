package content_type

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"study_savvy_api_go/api/response/utils"
)

func MiddleWareApplicationJson() gin.HandlerFunc {
	return func(c *gin.Context) {

		contentType := c.ContentType()

		if contentType != "application/json" {
			e := utils.Error{Error: "Content-Type must be application/json not " + contentType}
			c.JSON(http.StatusUnsupportedMediaType, e)
			c.Abort()
			return
		}

		c.Next()
	}
}
