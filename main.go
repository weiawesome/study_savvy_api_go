package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"study_savvy_api_go/api/middleware/request/ai_predict"
	"study_savvy_api_go/api/utils"
)

func main() {
	if err := utils.InitDB(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return
	}
	if err := utils.InitRedis(); err != nil {
		log.Fatalf("Failed to connect to the redis: %v", err)
		return
	}
	router := gin.Default()
	router.Use(ai_predict.MiddlewareAiPredictOcrContent())
	router.POST("/upload", uploadHandler)
	router.Run(":8080")
	//r := routes.InitRoutes()
	//err := r.Run()
	//if err != nil {
	//	return
	//}
	//log.Printf("Listening on port %s", port)
	//log.Printf("Open http://localhost:%s in the browser", port)
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func uploadHandler(c *gin.Context) {
	// 在這裡處理上傳的文件
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully!"})
}
