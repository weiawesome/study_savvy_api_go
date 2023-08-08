package routes

import (
	"github.com/gin-gonic/gin"
	"study_savvy_api_go/internal/repository/redis"
	"study_savvy_api_go/internal/repository/sql"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	err := r.SetTrustedProxies(nil)
	if err != nil {
		return nil
	}

	basicRouter := r.Group("/api")

	userRouter := basicRouter.Group("/User")
	nlpEditRouter := basicRouter.Group("/NLP_edit")
	accessMethodRouter := basicRouter.Group("/Access_method")
	mailRouter := basicRouter.Group("/Mail")
	filesRouter := basicRouter.Group("/Files")
	aiPredictRouter := basicRouter.Group("/Predict")
	oauthRouter := basicRouter.Group("/Oauth")
	informationRouter := basicRouter.Group("/Information")

	sqlRepository := sql.NewRepository()
	redisRepository := redis.NewRepository()

	InitUserRoutes(userRouter, sqlRepository, redisRepository)
	InitNlpEditRoutes(nlpEditRouter, sqlRepository, redisRepository)
	InitAccessMethodRoutes(accessMethodRouter, sqlRepository, redisRepository)
	InitMailRoutes(mailRouter, sqlRepository, redisRepository)
	InitFilesRoutes(filesRouter, sqlRepository, redisRepository)
	InitAiPredictRoutes(aiPredictRouter, sqlRepository, redisRepository)
	InitOauthRoutes(oauthRouter, sqlRepository, redisRepository)
	InitInformationRoutes(informationRouter, sqlRepository, redisRepository)

	return r
}
