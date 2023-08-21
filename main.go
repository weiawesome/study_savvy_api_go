package main

import (
	"github.com/gin-gonic/gin"
	"study_savvy_api_go/api/routes"
	"study_savvy_api_go/api/utils"
)

func main() {
	utils.InitLogger()

	if err := utils.InitDB(); err != nil {
		utils.LogFatal(utils.LogData{Event: "Failed to connect to the database", User: "system", Details: err.Error()})
		return
	}
	defer func() {
		err := utils.CloseDB()
		if err != nil {
			utils.LogFatal(utils.LogData{Event: "Failed to disconnect to the database", User: "system", Details: err.Error()})
		}
	}()

	if err := utils.InitRedis(); err != nil {
		utils.LogFatal(utils.LogData{Event: "Failed to connect to the redis", User: "system", Details: err.Error()})
		return
	}
	defer func() {
		err := utils.CloseRedis()
		if err != nil {
			utils.LogFatal(utils.LogData{Event: "Failed to disconnect to the redis", User: "system", Details: err.Error()})
		}
	}()

	gin.SetMode(gin.ReleaseMode)
	r := routes.InitRoutes()
	err := r.Run()

	if err != nil {
		return
	}
}
