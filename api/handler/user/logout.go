package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	responseUtils "study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/api/utils"
	"study_savvy_api_go/internal/service/user"
	"time"
)

type HandlerLogout struct {
	Service user.ServiceLogout
}

func (h *HandlerLogout) Handle(c *gin.Context) {
	jti, ok := c.Get("jti")
	if !ok {
		go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Data not found in context"})
		e := responseUtils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}

	if stringData, ok := jti.(string); ok {
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

			go utils.LogInfo(utils.LogData{Event: "Success request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header})

			c.JSON(http.StatusCreated, result)
		} else {
			go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: err.Error()}
			c.JSON(http.StatusInternalServerError, e)
		}
	} else {
		go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Type Assertion error"})
		e := responseUtils.Error{Error: "Internal error"}
		c.JSON(http.StatusInternalServerError, e)
	}

}
