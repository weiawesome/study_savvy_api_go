package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"study_savvy_api_go/api/response/files"
	"study_savvy_api_go/api/routes"
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
	//fmt.Println("")
	//sqlRepository := sql.NewRepository()
	//a, b, c := sqlRepository.ReadFileByPage("wei891013@gmail.com", 1, 10)
	//if c != nil {
	//	fmt.Println(c)
	//}
	//fmt.Println(a)
	//fmt.Println(b)
	//file := model.File{Id: uuid.New().String(), UserMail: "wei891013@gmail.com", Type: "ASR", Status: "SUCCESS", CreatedAt: time.Now()}
	//e := sqlRepository.CreateFile(file)
	//fmt.Println(e)

	r := routes.InitRoutes()
	err := r.Run()
	if err != nil {
		return
	}
	//log.Printf("Listening on port %s", port)
	//log.Printf("Open http://localhost:%s in the browser", port)
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
func test(c *gin.Context) {
	result := files.AudioFile{FilePath: "A.txt"}
	if err := result.Exist(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	err := result.CanOpenAndSent(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}
