package jwt

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	responseUtils "study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/api/utils"
	"study_savvy_api_go/internal/repository/redis"
)

func validateBlackList(jwtToken string, repository redis.Repository) error {
	jti := utils.InformationJwt(jwtToken).ID
	return repository.ValidateInBlacklist(jti)
}
func validateContent(jwtToken string) error {
	return utils.ValidateJwt(jwtToken)
}
func validateCookie(jwtToken string, csrfToken string) error {
	return utils.ValidateJwtCsrf(jwtToken, csrfToken)
}
func (m *MiddlewareJwt) JwtSecure() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtTokenCookie, err := c.Cookie("access_token_cookie")
		csrfToken := c.GetHeader("X-CSRF-TOKEN")

		jwtTokenContent := c.GetHeader("Authorization")
		if strings.HasPrefix(jwtTokenContent, "Bearer ") {
			jwtTokenContent, _ = strings.CutPrefix(jwtTokenContent, "Bearer ")
		} else {
			jwtTokenContent = ""
		}

		statusCookie := true
		statusContent := true

		if err != nil {
			statusCookie = false
		} else {
			statusCookie = statusCookie && (validateBlackList(jwtTokenCookie, m.repository) == nil) && (validateCookie(jwtTokenCookie, csrfToken) == nil)
		}
		statusContent = statusContent && (validateBlackList(jwtTokenContent, m.repository) == nil) && (validateContent(jwtTokenContent) == nil)

		if statusCookie {
			c.Set("jwt", jwtTokenCookie)
			c.Next()
		} else if statusContent {
			c.Set("jwt", jwtTokenContent)
			c.Next()
		} else {
			e := responseUtils.Error{Error: "Invalidate jwtToken"}
			c.JSON(http.StatusUnprocessableEntity, e)
			c.Abort()
			return
		}

	}
}
