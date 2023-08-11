package routes

import (
	"github.com/gin-gonic/gin"
	handlerUser "study_savvy_api_go/api/handler/user"
	"study_savvy_api_go/api/middleware/content_type"
	"study_savvy_api_go/api/middleware/jwt"
	requestUser "study_savvy_api_go/api/middleware/request/user"
	"study_savvy_api_go/internal/repository/redis"
	"study_savvy_api_go/internal/repository/sql"
	"study_savvy_api_go/internal/service/user"
)

func InitUserRoutes(r *gin.RouterGroup, sqlRepository *sql.Repository, redisRepository *redis.Repository) {
	middlewareJwt := jwt.MiddlewareJwt{Repository: redisRepository}
	r.POST("/login/app", content_type.MiddleWareApplicationJson(), requestUser.MiddleWareLoginAppContent(), (&handlerUser.HandlerLoginApp{Service: user.ServiceLoginApp{Repository: *sqlRepository}}).Handle)
	r.POST("/login/web", content_type.MiddleWareApplicationJson(), requestUser.MiddleWareLoginWebContent(), (&handlerUser.HandlerLoginWeb{Service: user.ServiceLoginWeb{Repository: *sqlRepository}}).Handle)
	r.DELETE("/logout", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), (&handlerUser.HandlerLogout{Service: user.ServiceLogout{Repository: *redisRepository}}).Handle)
	r.POST("/signup", content_type.MiddleWareApplicationJson(), requestUser.MiddleWareSignupContent(), (&handlerUser.HandlerSignup{Service: user.ServiceSignup{Repository: *sqlRepository}}).Handle)
}
