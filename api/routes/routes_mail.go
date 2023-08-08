package routes

import (
	"github.com/gin-gonic/gin"
	handlerMail "study_savvy_api_go/api/handler/mail"
	requestMail "study_savvy_api_go/api/middleware/request/mail"
	"study_savvy_api_go/internal/repository/redis"
	"study_savvy_api_go/internal/repository/sql"
	"study_savvy_api_go/internal/service/mail"
)

func InitMailRoutes(r *gin.RouterGroup, sqlRepository *sql.Repository, redisRepository *redis.Repository) {
	r.POST("/verification", requestMail.MiddlewareMailVerificationContent(), (&handlerMail.HandlerMailVerification{Service: mail.ServiceMailVerification{RedisRepository: *redisRepository}}).Handle)
	r.GET("/verification/:mail/:code", requestMail.MiddleWareMailVerifyContent(), (&handlerMail.HandlerMailVerify{Service: mail.ServiceMailVerify{RedisRepository: *redisRepository}}).Handle)
}
