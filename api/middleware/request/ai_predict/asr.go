package ai_predict

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"study_savvy_api_go/api/request/ai_predict"
)

func MiddlewareAiPredictAsrContent() gin.HandlerFunc {
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
		prompt := c.PostForm("prompt")

		uploadDir := "uploads"
		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			err := os.Mkdir(uploadDir, os.ModePerm)
			if err != nil {
				return
			}
		}

		data := ai_predict.Asr{File: handler, Prompt: prompt}
		fileType, err := data.Validate()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		id := uuid.New().String()
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

		data.File.Filename = filePath
		c.Set("data", data)

		c.Next()
	}
}
