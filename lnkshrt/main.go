package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"larssonoliver.com/lnkshrt/app"
	"larssonoliver.com/lnkshrt/app/config"
)

func logInit() {
	log.SetPrefix("\033[35m[" + config.Executable() + "]:\033[0m ")
}

func corsSetup(router *mux.Router) http.Handler {
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	ttl := handlers.MaxAge(3600)
	origins := handlers.AllowedOrigins(config.Origins())
	return handlers.CORS(origins, methods, ttl)(router)
}

func main() {
	logInit()
	app := app.New()

	port := fmt.Sprintf(":%d", config.Port())

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, corsSetup(app.Router)))
}
