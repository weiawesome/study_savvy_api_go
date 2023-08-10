package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	responseUtils "study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/api/utils"
	"study_savvy_api_go/internal/repository/redis"
)

func validateBlackList(jwtToken string, repository *redis.Repository) error {
	jwt := utils.InformationJwt(jwtToken)
	if jwt != nil {
		return errors.New("JwtToken error")
	}
	jti := jwt.ID
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
		statusCookie := true
		statusContent := true

		jwtTokenCookie, err := c.Cookie("access_token_cookie")
		csrfToken := c.GetHeader("X-CSRF-TOKEN")

		jwtTokenContent := c.GetHeader("Authorization")
		if strings.HasPrefix(jwtTokenContent, "Bearer ") {
			jwtTokenContent, _ = strings.CutPrefix(jwtTokenContent, "Bearer ")
			statusContent = statusContent && (validateBlackList(jwtTokenContent, m.Repository) == nil) && (validateContent(jwtTokenContent) == nil)
		} else {
			statusContent = false
		}

		if err == nil {
			statusCookie = statusCookie && (validateBlackList(jwtTokenCookie, m.Repository) == nil) && (validateCookie(jwtTokenCookie, csrfToken) == nil)
		} else {
			statusCookie = false
		}

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
