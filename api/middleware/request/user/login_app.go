package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"study_savvy_api_go/api/request/user"
)

func LoginAppContentMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Content-Type") != "application/json" {
			c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "Content-Type must be application/json"})
			c.Abort()
			return
		}

		var data user.LoginApp
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
			c.Abort()
			return
		}
		if data.Mail == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Mail can't be empty"})
			c.Abort()
			return
		}
		if data.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Password can't be empty"})
			c.Abort()
			return
		}

		c.Set("data", data)

		c.Next()
	}
}
