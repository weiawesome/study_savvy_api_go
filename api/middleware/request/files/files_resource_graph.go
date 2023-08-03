package files

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"study_savvy_api_go/api/response/utils"
)

func MiddleWareFilesResourceGraphContent() gin.HandlerFunc {
	return func(c *gin.Context) {
		Id := c.Param("file_id")
		IdUuid, err := uuid.Parse(Id)

		if err != nil {
			e := utils.Error{Error: "Parameter error not uuid"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		c.Set("id", IdUuid.String())
		c.Next()
	}
}
