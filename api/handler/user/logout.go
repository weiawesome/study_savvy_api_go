package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/internal/service/user"
	"time"
)

type LogoutHandler struct {
	Service user.LogoutService
}

func (h *LogoutHandler) Handle(c *gin.Context) {
	jwt, ok := c.Get("jwt")
	if !ok {
		e := utils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	if stringData, ok := jwt.(string); ok {
		result, err := h.Service.Logout(stringData)
		if err == nil {
			cookieJwt := &http.Cookie{
				Name:     "access_token_cookie",
				Value:    "",
				MaxAge:   -1,
				Expires:  time.Now(),
				Secure:   true,
				HttpOnly: true,
			}

			cookieCsrf := &http.Cookie{
				Name:    "csrf_access_token",
				Value:   "",
				MaxAge:  -1,
				Expires: time.Now(),
				Secure:  true,
			}
			c.SetCookie(cookieJwt.Name, cookieJwt.Value, cookieJwt.MaxAge, cookieJwt.Path, cookieJwt.Domain, cookieJwt.Secure, cookieJwt.HttpOnly)
			c.SetCookie(cookieCsrf.Name, cookieCsrf.Value, cookieCsrf.MaxAge, cookieCsrf.Path, cookieCsrf.Domain, cookieCsrf.Secure, cookieCsrf.HttpOnly)
			c.JSON(http.StatusCreated, result)
		} else {
			e := utils.Error{Error: err.Error()}
			c.JSON(http.StatusInternalServerError, e)
		}
	} else {
		e := utils.Error{Error: "Internal error"}
		c.JSON(http.StatusInternalServerError, e)
	}

}
