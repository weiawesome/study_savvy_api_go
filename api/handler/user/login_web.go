package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	userRequest "study_savvy_api_go/api/request/user"
	"study_savvy_api_go/api/response/utils"
	"study_savvy_api_go/internal/service/user"
	"time"
)

type LoginWebHandler struct {
	Service user.LoginWebService
}

func (h *LoginWebHandler) Handle(c *gin.Context) {
	data, ok := c.Get("data")
	if !ok {
		e := utils.Error{Error: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}

	if jsonData, ok := data.(userRequest.LoginWeb); ok {
		result, resultToken, err := h.Service.Login(jsonData)
		if err == nil {
			expiredDays, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE_DAYS"))
			if expiredDays == 0 {
				expiredDays = 1
			}
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
			c.JSON(http.StatusOK, result)
		} else if errors.As(err, &utils.RegistrationError{}) {
			e := utils.Error{Error: err.Error()}
			c.JSON(http.StatusUnauthorized, e)
		} else {
			e := utils.Error{Error: err.Error()}
			c.JSON(http.StatusInternalServerError, e)
		}
	} else {
		e := utils.Error{Error: "Internal error"}
		c.JSON(http.StatusInternalServerError, e)
	}
}
