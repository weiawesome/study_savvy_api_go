package files

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	responseUtils "study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/api/utils"
)

func MiddleWareFilesSpecificDeleteContent() gin.HandlerFunc {
	return func(c *gin.Context) {
		Id := c.Param("file_id")
		IdUuid, err := uuid.Parse(Id)

		if err != nil {
			go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: "Parameter error not uuid"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		c.Set("id", IdUuid.String())
		c.Next()
	}
}
