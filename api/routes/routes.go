package routes

import (
	"github.com/gin-gonic/gin"
	"study_savvy_api_go/internal/repository/redis"
	"study_savvy_api_go/internal/repository/sql"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	basicRouter := r.Group("/api")

	userRouter := basicRouter.Group("")
	nlpEditRouter := basicRouter.Group("/NLP_edit")
	accessMethodRouter := basicRouter.Group("/Access_method")
	mailRouter := basicRouter.Group("/verification")
	filesRouter := basicRouter.Group("/files")
	aiPredictRouter := basicRouter.Group("/predict")
	oauthRouter := basicRouter.Group("/oauth")
	informationRouter := basicRouter.Group("/information")

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
