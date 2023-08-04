package routes

import (
	"github.com/gin-gonic/gin"
	handlerAccessMethod "study_savvy_api_go/api/handler/access_method"
	handlerFiles "study_savvy_api_go/api/handler/files"
	handlerInformation "study_savvy_api_go/api/handler/information"
	handlerUser "study_savvy_api_go/api/handler/user"
	"study_savvy_api_go/api/middleware/jwt"
	requestAccessMethod "study_savvy_api_go/api/middleware/request/access_method"
	requestFiles "study_savvy_api_go/api/middleware/request/files"
	requestInformation "study_savvy_api_go/api/middleware/request/information"
	requestUser "study_savvy_api_go/api/middleware/request/user"
	"study_savvy_api_go/internal/repository/redis"
	"study_savvy_api_go/internal/repository/sql"
	"study_savvy_api_go/internal/service/access_method"
	"study_savvy_api_go/internal/service/files"
	"study_savvy_api_go/internal/service/information"
	"study_savvy_api_go/internal/service/user"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/api")
	//nlpEditRouter := r.Group("/api/NLP_edit")
	accessMethodRouter := r.Group("/api/Access_method")
	//mailRouter := r.Group("/api/verification")
	filesRouter := r.Group("/api/files")
	//aiPredictRouter := r.Group("/api/predict")
	//oauthRouter := r.Group("/api/oauth")
	informationRouter := r.Group("/api/information")

	sqlRepository := sql.NewRepository()
	redisRepository := redis.NewRepository()

	userRouter.POST("/login/app", requestUser.MiddleWareLoginContent(), (&handlerUser.HandlerLoginApp{Service: user.ServiceLoginApp{Repository: *sqlRepository}}).Handle)
	userRouter.POST("/login/web", requestUser.MiddleWareLoginContent(), (&handlerUser.HandlerLoginWeb{Service: user.ServiceLoginWeb{Repository: *sqlRepository}}).Handle)
	userRouter.DELETE("/logout", jwt.MiddlewareJwtSecure(), (&handlerUser.HandlerLogout{Service: user.ServiceLogout{Repository: *redisRepository}}).Handle)
	userRouter.POST("/signup", requestUser.MiddleWareSignupContent(), (&handlerUser.HandlerSignup{Service: user.ServiceSignup{Repository: *sqlRepository}}).Handle)
	//
	//nlpEditRouter.PUT("/ASR/{file_id}", AuthHomeHandler)
	//nlpEditRouter.PUT("/OCR/{file_id}", AuthHomeHandler)
	//
	accessMethodRouter.PUT("/access_token", jwt.MiddlewareJwtSecure(), jwt.MiddlewareJwtInformation(), requestAccessMethod.MiddleWareAccessTokenEditContent(), (&handlerAccessMethod.HandlerAccessMethodAccessToken{Service: access_method.ServiceAccessMethodAccessToken{Repository: *sqlRepository}}).Handle)
	accessMethodRouter.PUT("/api_key", jwt.MiddlewareJwtSecure(), jwt.MiddlewareJwtInformation(), requestAccessMethod.MiddleWareApiKeyEditContent(), (&handlerAccessMethod.HandlerAccessMethodApiKey{Service: access_method.ServiceAccessMethodApiKey{Repository: *sqlRepository}}).Handle)

	//
	//mailRouter.POST("/", AuthHomeHandler)
	//mailRouter.GET("/{mail}/{code}", AuthHomeHandler)
	//
	filesRouter.GET("", jwt.MiddlewareJwtSecure(), jwt.MiddlewareJwtInformation(), requestFiles.MiddleWareFilesContent(), (&handlerFiles.HandlerFiles{Service: files.ServiceFiles{Repository: *sqlRepository}}).Handle)
	filesRouter.GET("/ASR", jwt.MiddlewareJwtSecure(), jwt.MiddlewareJwtInformation(), requestFiles.MiddleWareFilesAsrContent(), (&handlerFiles.HandlerFilesAsr{Service: files.ServiceFilesAsr{Repository: *sqlRepository}}).Handle)
	filesRouter.GET("/OCR", jwt.MiddlewareJwtSecure(), jwt.MiddlewareJwtInformation(), requestFiles.MiddleWareFilesOcrContent(), (&handlerFiles.HandlerFilesOcr{Service: files.ServiceFilesOcr{Repository: *sqlRepository}}).Handle)
	filesRouter.GET("/resources/audio/:file_id", jwt.MiddlewareJwtSecure(), jwt.MiddlewareJwtInformation(), requestFiles.MiddleWareFilesResourceAudioContent(), (&handlerFiles.HandlerFilesResourceAudio{Service: files.ServiceFilesResourceAudio{Repository: *sqlRepository}}).Handle)
	filesRouter.GET("/resources/graph/:file_id", jwt.MiddlewareJwtSecure(), jwt.MiddlewareJwtInformation(), requestFiles.MiddleWareFilesResourceGraphContent(), (&handlerFiles.HandlerFilesResourceGraph{Service: files.ServiceFilesResourceGraph{Repository: *sqlRepository}}).Handle)
	filesRouter.GET("/:file_id", jwt.MiddlewareJwtSecure(), jwt.MiddlewareJwtInformation(), requestFiles.MiddleWareFilesSpecificContent(), (&handlerFiles.HandlerFilesSpecific{Service: files.ServiceFilesSpecific{Repository: *sqlRepository}}).Handle)
	filesRouter.DELETE("/:file_id", jwt.MiddlewareJwtSecure(), jwt.MiddlewareJwtInformation(), requestFiles.MiddleWareFilesSpecificDeleteContent(), (&handlerFiles.HandlerFilesSpecificDelete{Service: files.ServiceFilesSpecificDelete{Repository: *sqlRepository}}).Handle)
	//
	//aiPredictRouter.POST("/ASR", AuthHomeHandler)
	//aiPredictRouter.POST("/OCR", AuthHomeHandler)
	//aiPredictRouter.POST("/OCR/text", AuthHomeHandler)
	//
	//oauthRouter.GET("/app/google", AuthHomeHandler)
	//oauthRouter.GET("/web/google", AuthHomeHandler)
	//oauthRouter.GET("/authorize/google", AuthHomeHandler)
	//
	informationRouter.GET("", jwt.MiddlewareJwtSecure(), jwt.MiddlewareJwtInformation(), (&handlerInformation.HandlerInformation{Service: information.ServiceInformation{Repository: *sqlRepository}}).Handle)
	informationRouter.PUT("", jwt.MiddlewareJwtSecure(), jwt.MiddlewareJwtInformation(), requestInformation.MiddleWareInformationEditContent(), (&handlerInformation.HandlerInformationEdit{Service: information.ServiceInformationEdit{Repository: *sqlRepository}}).Handle)
	informationRouter.PUT("/password_edit", jwt.MiddlewareJwtSecure(), jwt.MiddlewareJwtInformation(), requestInformation.MiddleWarePasswordEditContent(), (&handlerInformation.HandlerPasswordEdit{Service: information.ServicePasswordEdit{Repository: *sqlRepository}}).Handle)

	return r
}
