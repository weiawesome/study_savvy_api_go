package jwtSecure

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	responseUtils "study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/api/utils"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken, _ := c.Cookie("access_token_cookie")
		csrfToken := c.GetHeader("X-CSRF-TOKEN")
		if status, err := utils.ValidateJwtCsrf(jwtToken, csrfToken); err != nil || status == false {
			jwt := c.GetHeader("Authorization")
			if strings.HasPrefix(jwt, "Bearer ") {
				jwt, _ = strings.CutPrefix(jwt, "Bearer ")
			} else {
				jwt = ""
			}
			if status, err := utils.ValidateJwt(jwt); err != nil || status == false {
				if status == false {
					e := responseUtils.Error{Error: "Invalidate jwtToken"}
					c.JSON(http.StatusUnprocessableEntity, e)
					c.Abort()
					return
				} else {
					e := responseUtils.Error{Error: "Internal error"}
					c.JSON(http.StatusInternalServerError, e)
					c.Abort()
					return
				}
			} else {
				c.Set("jwt", jwt)
				c.Next()
			}
		} else {
			c.Set("jwt", jwtToken)
			c.Next()
		}

	}
}
