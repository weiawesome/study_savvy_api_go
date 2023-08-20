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
	responseUtils "study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/api/utils"
)

func MiddlewareAiPredictOcrContent() gin.HandlerFunc {
	return func(c *gin.Context) {

		file, handler, err := c.Request.FormFile("file")

		if err != nil {
			go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: "Failed to get file from form"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: err.Error()})
				e := responseUtils.Error{Error: "Failed to close file"}
				c.JSON(http.StatusBadRequest, e)
				c.Abort()
			}
		}(file)
		prompt := c.PostForm("prompt")

		uploadDir := utils.EnvGraphDirectory()
		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			err := os.Mkdir(uploadDir, os.ModePerm)
			if err != nil {
				go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: err.Error()})
				return
			}
		}

		data := ai_predict.Ocr{File: handler, Prompt: prompt}
		fileType, err := data.Validate()
		if err != nil {
			go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: err.Error()}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		id := uuid.New().String()
		filePath := fmt.Sprintf("%s/%s", uploadDir, id+fileType)
		outFile, err := os.Create(filePath)
		if err != nil {
			go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: "Failed to create file on server"}
			c.JSON(http.StatusInternalServerError, e)
			c.Abort()
			return
		}
		defer func(outFile *os.File) {
			err := outFile.Close()
			if err != nil {
				go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: err.Error()})
				e := responseUtils.Error{Error: "Failed to create file"}
				c.JSON(http.StatusBadRequest, e)
				c.Abort()
				return
			}
		}(outFile)

		_, err = file.Seek(0, 0)
		if err != nil {
			go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: "Failed to rewind file"}
			c.JSON(http.StatusInternalServerError, e)
			c.Abort()
			return
		}

		_, err = io.Copy(outFile, file)
		if err != nil {
			go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: "Failed to write file on server"}
			c.JSON(http.StatusInternalServerError, e)
			c.Abort()
			return
		}

		data.File.Filename = filePath
		c.Set("data", data)

		c.Next()
	}
}
