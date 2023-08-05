package ai_predict

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func ValidateAiPredictOcr(fileType string) (string, bool) {
	supportedTypes := []string{".jpg", ".jpeg", ".png"}
	for _, t := range supportedTypes {
		if strings.HasSuffix(fileType, t) {
			return t, true
		}
	}
	return "", false
}

func MiddlewareAiPredictOcrContent() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "POST" || c.ContentType() != "multipart/form-data" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Only multipart/form-data POST requests are allowed"})
			return
		}

		file, handler, err := c.Request.FormFile("file")

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to get file from form"})
			return
		}
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to close file"})
			}
		}(file)

		uploadDir := "uploads"
		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			err := os.Mkdir(uploadDir, os.ModePerm)
			if err != nil {
				return
			}
		}

		fileType, status := ValidateAiPredictOcr(handler.Filename)
		if status == false {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "File type is not support"})
			return
		}

		id := uuid.New().String()
		c.Set("id", id)

		filePath := fmt.Sprintf("%s/%s", uploadDir, id+fileType)
		outFile, err := os.Create(filePath)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create file on server"})
			return
		}
		defer func(outFile *os.File) {
			err := outFile.Close()
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to create file"})
				return
			}
		}(outFile)

		_, err = file.Seek(0, 0)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to rewind file"})
			return
		}

		_, err = io.Copy(outFile, file)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to write file on server"})
			return
		}
		fmt.Println(filePath)

		prompt := c.PostForm("prompt")
		c.Set("prompt", prompt)

		c.Next()
	}
}
