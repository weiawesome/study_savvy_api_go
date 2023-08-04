package routes

import (
	"github.com/gin-gonic/gin"
	"study_savvy_api_go/internal/repository/redis"
	"study_savvy_api_go/internal/repository/sql"
)

func InitNlpEditRoutes(r *gin.RouterGroup, sqlRepository *sql.Repository, redisRepository *redis.Repository) {
	//middlewareJwt := jwt.MiddlewareJwt{Repository: redisRepository}
	//r.PUT("/ASR/:file_id", AuthHomeHandler)
	//r.PUT("/OCR/:file_id", AuthHomeHandler)
}
