package routes

import (
	"github.com/gin-gonic/gin"
	handlerFiles "study_savvy_api_go/api/handler/files"
	"study_savvy_api_go/api/middleware/jwt"
	requestFiles "study_savvy_api_go/api/middleware/request/files"
	"study_savvy_api_go/internal/repository/redis"
	"study_savvy_api_go/internal/repository/sql"
	"study_savvy_api_go/internal/service/files"
	"study_savvy_api_go/internal/service/logger"
)

func InitFilesRoutes(r *gin.RouterGroup, sqlRepository *sql.Repository, redisRepository *redis.Repository) {
	middlewareJwt := jwt.MiddlewareJwt{Repository: redisRepository}
	r.GET("", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), requestFiles.MiddleWareFilesContent(), (&handlerFiles.HandlerFiles{Service: files.ServiceFiles{Repository: *sqlRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).Handle)
	r.GET("/ASR", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), requestFiles.MiddleWareFilesAsrContent(), (&handlerFiles.HandlerFilesAsr{Service: files.ServiceFilesAsr{Repository: *sqlRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).Handle)
	r.GET("/OCR", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), requestFiles.MiddleWareFilesOcrContent(), (&handlerFiles.HandlerFilesOcr{Service: files.ServiceFilesOcr{Repository: *sqlRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).Handle)
	r.GET("/resources/audio/:file_id", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), requestFiles.MiddleWareFilesResourceAudioContent(), (&handlerFiles.HandlerFilesResourceAudio{Service: files.ServiceFilesResourceAudio{Repository: *sqlRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).Handle)
	r.GET("/resources/graph/:file_id", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), requestFiles.MiddleWareFilesResourceGraphContent(), (&handlerFiles.HandlerFilesResourceGraph{Service: files.ServiceFilesResourceGraph{Repository: *sqlRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).Handle)
	r.GET("/:file_id", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), requestFiles.MiddleWareFilesSpecificContent(), (&handlerFiles.HandlerFilesSpecific{Service: files.ServiceFilesSpecific{Repository: *sqlRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).Handle)
	r.DELETE("/:file_id", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), requestFiles.MiddleWareFilesSpecificDeleteContent(), (&handlerFiles.HandlerFilesSpecificDelete{Service: files.ServiceFilesSpecificDelete{Repository: *sqlRepository}, LogService: logger.ServiceLogger{Repository: *redisRepository}}).Handle)
}
