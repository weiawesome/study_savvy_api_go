package routes

import (
	"github.com/gin-gonic/gin"
	"study_savvy_api_go/internal/repository/redis"
	"study_savvy_api_go/internal/repository/sql"
)

func InitAiPredictRoutes(r *gin.RouterGroup, sqlRepository *sql.Repository, redisRepository *redis.Repository) {
	//middlewareJwt := jwt.MiddlewareJwt{Repository: redisRepository}
	//r.POST("/ASR", AuthHomeHandler)
	//r.POST("/OCR", AuthHomeHandler)
	//r.POST("/OCR/text", AuthHomeHandler)
}
