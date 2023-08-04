package main

import (
	"log"
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

	r := routes.InitRoutes()
	err := r.Run()
	if err != nil {
		return
	}
	//log.Printf("Listening on port %s", port)
	//log.Printf("Open http://localhost:%s in the browser", port)
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
