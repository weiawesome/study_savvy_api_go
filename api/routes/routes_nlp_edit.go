package routes

import (
	"github.com/gin-gonic/gin"
	handlerNlpEdit "study_savvy_api_go/api/handler/nlp_edit"
	"study_savvy_api_go/api/middleware/content_type"
	"study_savvy_api_go/api/middleware/jwt"
	requestNlpEdit "study_savvy_api_go/api/middleware/request/nlp_edit"
	"study_savvy_api_go/internal/repository/redis"
	"study_savvy_api_go/internal/repository/sql"
	"study_savvy_api_go/internal/service/nlp_edit"
)

func InitNlpEditRoutes(r *gin.RouterGroup, sqlRepository *sql.Repository, redisRepository *redis.Repository) {
	middlewareJwt := jwt.MiddlewareJwt{Repository: redisRepository}
	r.PUT("/ASR/:file_id", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), content_type.MiddleWareApplicationJson(), requestNlpEdit.MiddlewareNlpEditAsrContent(), (&handlerNlpEdit.HandlerNlpEditAsr{Service: nlp_edit.ServiceNlpEditAsr{SqlRepository: *sqlRepository, RedisRepository: *redisRepository}}).Handle)
	r.PUT("/OCR/:file_id", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), content_type.MiddleWareApplicationJson(), requestNlpEdit.MiddlewareNlpEditOcrContent(), (&handlerNlpEdit.HandlerNlpEditOcr{Service: nlp_edit.ServiceNlpEditOcr{SqlRepository: *sqlRepository, RedisRepository: *redisRepository}}).Handle)
}
