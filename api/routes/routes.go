package routes

import (
	"github.com/gin-gonic/gin"
	handlerUser "study_savvy_api_go/api/handler/user"
	requestUser "study_savvy_api_go/api/middleware/request/user"
	"study_savvy_api_go/internal/repository"
	"study_savvy_api_go/internal/service/user"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/api")
	//nlpEditRouter := r.Group("/api/NLP_edit")
	//accessMethodRouter := r.Group("/api/Access_method")
	//mailRouter := r.Group("/api/verification")
	//filesRouter := r.Group("/api/files")
	//aiPredictRouter := r.Group("/api/predict")
	//oauthRouter := r.Group("/api/oauth")
	//informationRouter := r.Group("/api/information")
	Repository := repository.NewRepository()
	userRouter.POST("/login/app", requestUser.LoginAppContentMiddleWare(), (&handlerUser.LoginAppHandler{Service: user.LoginAppService{Repository: *Repository}}).Handle)
	//userRouter.POST("/login/web", AuthHomeHandler)
	//userRouter.DELETE("/logout", AuthHomeHandler)
	userRouter.POST("/signup", requestUser.SignupContentMiddleWare(), (&handlerUser.SignupHandler{Service: user.SignupService{Repository: *Repository}}).Handle)
	//
	//nlpEditRouter.PUT("/ASR/{file_id}", AuthHomeHandler)
	//nlpEditRouter.PUT("/OCR/{file_id}", AuthHomeHandler)
	//
	//accessMethodRouter.PUT("/access_token", AuthHomeHandler)
	//accessMethodRouter.PUT("/api_key", AuthHomeHandler)
	//
	//mailRouter.POST("/", AuthHomeHandler)
	//mailRouter.GET("/{mail}/{code}", AuthHomeHandler)
	//
	//filesRouter.GET("/", AuthHomeHandler)
	//filesRouter.GET("/ASR", AuthHomeHandler)
	//filesRouter.GET("/OCR", AuthHomeHandler)
	//filesRouter.GET("/resources/audio/{file_id}", AuthHomeHandler)
	//filesRouter.GET("/resources/graph/{file_id}", AuthHomeHandler)
	//filesRouter.GET("/{file_id}", AuthHomeHandler)
	//filesRouter.DELETE("/{file_id}", AuthHomeHandler)
	//
	//aiPredictRouter.POST("/ASR", AuthHomeHandler)
	//aiPredictRouter.POST("/OCR", AuthHomeHandler)
	//aiPredictRouter.POST("/OCR/text", AuthHomeHandler)
	//
	//oauthRouter.GET("/app/google", AuthHomeHandler)
	//oauthRouter.GET("/web/google", AuthHomeHandler)
	//oauthRouter.GET("/authorize/google", AuthHomeHandler)
	//
	//informationRouter.GET("/", AuthHomeHandler)
	//informationRouter.PUT("/", AuthHomeHandler)
	//informationRouter.PUT("/password_edit", AuthHomeHandler)

	return r
}
