package jwt

import (
	"github.com/gin-gonic/gin"
	"net/http"
	responseUtils "study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/api/utils"
)

func (m *MiddlewareJwt) JwtInformation() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt, ok := c.Get("jwt")
		if !ok {
			e := responseUtils.Error{Error: "Data not found in context"}
			c.JSON(http.StatusInternalServerError, e)
			return
		}
		if stringData, ok := jwt.(string); ok {
			jwtToken := utils.InformationJwt(stringData)
			if jwtToken == nil {
				e := responseUtils.Error{Error: "Jwt can't parse"}
				c.JSON(http.StatusUnprocessableEntity, e)
				return
			} else {
				c.Set("user", jwtToken.Subject)
				c.Set("jti", jwtToken.ID)
				c.Next()
			}
		} else {
			e := responseUtils.Error{Error: "Data can't parse"}
			c.JSON(http.StatusInternalServerError, e)
			return
		}
	}
}
