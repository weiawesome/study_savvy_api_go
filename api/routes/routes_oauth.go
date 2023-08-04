package routes

import (
	"github.com/gin-gonic/gin"
	"study_savvy_api_go/internal/repository/redis"
	"study_savvy_api_go/internal/repository/sql"
)

func InitOauthRoutes(r *gin.RouterGroup, sqlRepository *sql.Repository, redisRepository *redis.Repository) {
	//middlewareJwt := jwt.MiddlewareJwt{Repository: redisRepository}
	//r.GET("/app/google", AuthHomeHandler)
	//r.GET("/web/google", AuthHomeHandler)
	//r.GET("/authorize/google", AuthHomeHandler)
}
