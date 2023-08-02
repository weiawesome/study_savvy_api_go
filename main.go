package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"study_savvy_api_go/api/routes"
	"study_savvy_api_go/api/utils"
)

type jwtClaim struct {
	Type  string `json:"type"`
	Fresh bool   `json:"fresh"`
	Csrf  string `json:"csrf"`
	jwt.RegisteredClaims
}

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
	//http.HandleFunc("/", indexHandler)
	//port := os.Getenv("PORT")
	//if port == "" {
	//	port = "8080"
	//	log.Printf("Defaulting to port %s", port)
	//}
	//
	//log.Printf("Listening on port %s", port)
	//log.Printf("Open http://localhost:%s in the browser", port)
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
