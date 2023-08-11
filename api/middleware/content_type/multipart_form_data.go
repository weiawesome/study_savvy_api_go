package content_type

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MiddleWareMultipartFormData() gin.HandlerFunc {
	return func(c *gin.Context) {

		contentType := c.ContentType()

		if contentType != "multipart/form-data" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Content-Type not " + contentType})
			return
		}

		c.Next()
	}
}
