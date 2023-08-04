package routes

import (
	"github.com/gin-gonic/gin"
	handlerFiles "study_savvy_api_go/api/handler/files"
	"study_savvy_api_go/api/middleware/jwt"
	requestFiles "study_savvy_api_go/api/middleware/request/files"
	"study_savvy_api_go/internal/repository/redis"
	"study_savvy_api_go/internal/repository/sql"
	"study_savvy_api_go/internal/service/files"
)

func InitFilesRoutes(r *gin.RouterGroup, sqlRepository *sql.Repository, redisRepository *redis.Repository) {
	middlewareJwt := jwt.MiddlewareJwt{Repository: redisRepository}
	r.GET("", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), requestFiles.MiddleWareFilesContent(), (&handlerFiles.HandlerFiles{Service: files.ServiceFiles{Repository: *sqlRepository}}).Handle)
	r.GET("/ASR", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), requestFiles.MiddleWareFilesAsrContent(), (&handlerFiles.HandlerFilesAsr{Service: files.ServiceFilesAsr{Repository: *sqlRepository}}).Handle)
	r.GET("/OCR", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), requestFiles.MiddleWareFilesOcrContent(), (&handlerFiles.HandlerFilesOcr{Service: files.ServiceFilesOcr{Repository: *sqlRepository}}).Handle)
	r.GET("/resources/audio/:file_id", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), requestFiles.MiddleWareFilesResourceAudioContent(), (&handlerFiles.HandlerFilesResourceAudio{Service: files.ServiceFilesResourceAudio{Repository: *sqlRepository}}).Handle)
	r.GET("/resources/graph/:file_id", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), requestFiles.MiddleWareFilesResourceGraphContent(), (&handlerFiles.HandlerFilesResourceGraph{Service: files.ServiceFilesResourceGraph{Repository: *sqlRepository}}).Handle)
	r.GET("/:file_id", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), requestFiles.MiddleWareFilesSpecificContent(), (&handlerFiles.HandlerFilesSpecific{Service: files.ServiceFilesSpecific{Repository: *sqlRepository}}).Handle)
	r.DELETE("/:file_id", middlewareJwt.JwtSecure(), middlewareJwt.JwtInformation(), requestFiles.MiddleWareFilesSpecificDeleteContent(), (&handlerFiles.HandlerFilesSpecificDelete{Service: files.ServiceFilesSpecificDelete{Repository: *sqlRepository}}).Handle)
}
