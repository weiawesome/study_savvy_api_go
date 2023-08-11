package routes

import (
	"github.com/gin-gonic/gin"
	handlerInformation "study_savvy_api_go/api/handler/information"
	"study_savvy_api_go/api/middleware/content_type"
	"study_savvy_api_go/api/middleware/jwt"
	requestInformation "study_savvy_api_go/api/middleware/request/information"
	"study_savvy_api_go/internal/repository/redis"
	"study_savvy_api_go/internal/repository/sql"
	"study_savvy_api_go/internal/service/information"
)

func InitInformationRoutes(r *gin.RouterGroup, sqlRepository *sql.Repository, redisRepository *redis.Repository) {
	middlewareJwt := jwt.MiddlewareJwt{Repository: redisRepository}
	r.GET("", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), (&handlerInformation.HandlerInformation{Service: information.ServiceInformation{Repository: *sqlRepository}}).Handle)
	r.PUT("", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), content_type.MiddleWareApplicationJson(), requestInformation.MiddleWareInformationEditContent(), (&handlerInformation.HandlerInformationEdit{Service: information.ServiceInformationEdit{Repository: *sqlRepository}}).Handle)
	r.PUT("/password_edit", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), content_type.MiddleWareApplicationJson(), requestInformation.MiddleWarePasswordEditContent(), (&handlerInformation.HandlerPasswordEdit{Service: information.ServicePasswordEdit{Repository: *sqlRepository}}).Handle)
}
