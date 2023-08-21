package routes

import (
	"github.com/gin-gonic/gin"
	handlerAccessMethod "study_savvy_api_go/api/handler/access_method"
	"study_savvy_api_go/api/middleware/content_type"
	"study_savvy_api_go/api/middleware/jwt"
	requestAccessMethod "study_savvy_api_go/api/middleware/request/access_method"
	"study_savvy_api_go/internal/repository/redis"
	"study_savvy_api_go/internal/repository/sql"
	"study_savvy_api_go/internal/service/access_method"
	"study_savvy_api_go/internal/service/logger"
)

func InitAccessMethodRoutes(r *gin.RouterGroup, sqlRepository *sql.Repository, redisRepository *redis.Repository) {
	middlewareJwt := jwt.MiddlewareJwt{Repository: redisRepository}
	r.PUT("/access_token", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), content_type.MiddleWareApplicationJson(), requestAccessMethod.MiddleWareAccessTokenEditContent(), (&handlerAccessMethod.HandlerAccessMethodAccessToken{Service: access_method.ServiceAccessMethodAccessToken{Repository: *sqlRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).Handle)
	r.PUT("/api_key", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), content_type.MiddleWareApplicationJson(), requestAccessMethod.MiddleWareApiKeyEditContent(), (&handlerAccessMethod.HandlerAccessMethodApiKey{Service: access_method.ServiceAccessMethodApiKey{Repository: *sqlRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).Handle)
}
