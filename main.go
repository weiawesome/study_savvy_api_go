package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"study_savvy_api_go/api/routes"
	"study_savvy_api_go/api/utils"
)

func main() {
	utils.InitLogger()
	if err := utils.InitDB(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return
	}
	if err := utils.InitRedis(); err != nil {
		log.Fatalf("Failed to connect to the redis: %v", err)
		return
	}

	gin.SetMode(gin.ReleaseMode)
	r := routes.InitRoutes()
	err := r.Run()

	if err != nil {
		return
	}
}
