package routes

import (
	"github.com/gin-gonic/gin"
	handlerAiPredict "study_savvy_api_go/api/handler/ai_predict"
	"study_savvy_api_go/api/middleware/content_type"
	"study_savvy_api_go/api/middleware/jwt"
	requestAiPredict "study_savvy_api_go/api/middleware/request/ai_predict"
	"study_savvy_api_go/internal/repository/redis"
	"study_savvy_api_go/internal/repository/sql"
	"study_savvy_api_go/internal/service/ai_predict"
)

func InitAiPredictRoutes(r *gin.RouterGroup, sqlRepository *sql.Repository, redisRepository *redis.Repository) {
	middlewareJwt := jwt.MiddlewareJwt{Repository: redisRepository}
	r.POST("/ASR", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), content_type.MiddleWareMultipartFormData(), requestAiPredict.MiddlewareAiPredictAsrContent(), (&handlerAiPredict.HandlerAiPredictAsr{Service: ai_predict.ServiceAiPredictAsr{SqlRepository: *sqlRepository, RedisRepository: *redisRepository}}).Handle)
	r.POST("/OCR", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), content_type.MiddleWareMultipartFormData(), requestAiPredict.MiddlewareAiPredictOcrContent(), (&handlerAiPredict.HandlerAiPredictOcr{Service: ai_predict.ServiceAiPredictOcr{SqlRepository: *sqlRepository, RedisRepository: *redisRepository}}).Handle)
	r.POST("/OCR_Text", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), content_type.MiddleWareApplicationJson(), requestAiPredict.MiddlewareAiPredictOcrTextContent(), (&handlerAiPredict.HandlerAiPredictOcrText{Service: ai_predict.ServiceAiPredictOcrText{SqlRepository: *sqlRepository, RedisRepository: *redisRepository}}).Handle)
}
