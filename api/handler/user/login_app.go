package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	userRequest "study_savvy_api_go/api/request/user"
)

func LoginAppHandler(c *gin.Context) {
	data, ok := c.Get("data")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Data not found in context"})
		return
	}

	if jsonData, ok := data.(userRequest.LoginApp); ok {
		mail := jsonData.Mail
		password := jsonData.Password

		c.JSON(http.StatusOK, gin.H{"message": "JSON data received", "username": mail, "password": password})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Data type mismatch"})
	}
}
