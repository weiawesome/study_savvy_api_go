package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	userRequest "study_savvy_api_go/api/request/user"
	responseUtils "study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/api/utils"
	"study_savvy_api_go/internal/service/user"
	"time"
)

type HandlerLoginWeb struct {
	Service user.ServiceLoginWeb
}

func (h *HandlerLoginWeb) Handle(c *gin.Context) {
	data, ok := c.Get("data")
	if !ok {
		go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), Header: c.Request.Header, Details: "Data not found in context"})
		e := responseUtils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}

	if jsonData, ok := data.(userRequest.LoginWeb); ok {
		result, resultToken, err := h.Service.Login(jsonData)
		if err == nil {
			expiredDays, _ := strconv.Atoi(utils.EnvJwtExpireDays())

			cookieJwt := &http.Cookie{
				Name:     "access_token_cookie",
				Value:    resultToken.JwtToken,
				Expires:  time.Now().Add(time.Duration(expiredDays) * 24 * time.Hour),
				Secure:   true,
				HttpOnly: true,
			}
			cookieCsrf := &http.Cookie{
				Name:    "csrf_access_token",
				Value:   resultToken.CsrfToken,
				Expires: time.Now().Add(time.Duration(expiredDays) * 24 * time.Hour),
				Secure:  true,
			}

			c.SetCookie(cookieJwt.Name, cookieJwt.Value, cookieJwt.MaxAge, cookieJwt.Path, cookieJwt.Domain, cookieJwt.Secure, cookieJwt.HttpOnly)
			c.SetCookie(cookieCsrf.Name, cookieCsrf.Value, cookieCsrf.MaxAge, cookieCsrf.Path, cookieCsrf.Domain, cookieCsrf.Secure, cookieCsrf.HttpOnly)

			go utils.LogInfo(utils.LogData{Event: "Success request", Method: c.Request.Method, Path: c.FullPath(), User: jsonData.Mail, Header: c.Request.Header})

			c.JSON(http.StatusOK, result)
		} else if errors.As(err, &responseUtils.RegistrationError{}) {
			go utils.LogWarn(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: jsonData.Mail, Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: err.Error()}
			c.JSON(http.StatusUnauthorized, e)
		} else {
			go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: jsonData.Mail, Header: c.Request.Header, Details: err.Error()})
			e := responseUtils.Error{Error: err.Error()}
			c.JSON(http.StatusInternalServerError, e)
		}
	} else {
		go utils.LogError(utils.LogData{Event: "Failure request", Method: c.Request.Method, Path: c.FullPath(), User: jsonData.Mail, Header: c.Request.Header, Details: "Type Assertion error"})
		e := responseUtils.Error{Error: "Internal error"}
		c.JSON(http.StatusInternalServerError, e)
	}
}
