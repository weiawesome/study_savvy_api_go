package routes

import (
	"github.com/gin-gonic/gin"
	handlerAccessMethod "study_savvy_api_go/api/handler/access_method"
	"study_savvy_api_go/api/middleware/jwt"
	requestAccessMethod "study_savvy_api_go/api/middleware/request/access_method"
	"study_savvy_api_go/internal/repository/redis"
	"study_savvy_api_go/internal/repository/sql"
	"study_savvy_api_go/internal/service/access_method"
)

func InitAccessMethodRoutes(r *gin.RouterGroup, sqlRepository *sql.Repository, redisRepository *redis.Repository) {
	middlewareJwt := jwt.MiddlewareJwt{Repository: redisRepository}
	r.PUT("/access_token", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), requestAccessMethod.MiddleWareAccessTokenEditContent(), (&handlerAccessMethod.HandlerAccessMethodAccessToken{Service: access_method.ServiceAccessMethodAccessToken{Repository: *sqlRepository}}).Handle)
	r.PUT("/api_key", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), requestAccessMethod.MiddleWareApiKeyEditContent(), (&handlerAccessMethod.HandlerAccessMethodApiKey{Service: access_method.ServiceAccessMethodApiKey{Repository: *sqlRepository}}).Handle)
}
