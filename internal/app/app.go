package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/larssonoliver/lnkshrt/internal/config"
	"github.com/larssonoliver/lnkshrt/internal/db"
)

type App struct {
	Router   *mux.Router
	Database *db.Database
}

func New() *App {
	db, err := db.New(config.DBFile())
	if err != nil {
		log.Fatalln(err)
	}

	a := &App{
		Router:   mux.NewRouter().StrictSlash(true),
		Database: db,
	}

	a.initRoutes()
	return a
}

func (a *App) Run() {
	port := fmt.Sprintf(":%d", config.Port())

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, corsSetup(a.Router)))
}

func corsSetup(router *mux.Router) http.Handler {
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	ttl := handlers.MaxAge(3600)
	origins := handlers.AllowedOrigins(config.Origins())
	return handlers.CORS(origins, methods, ttl)(router)
}
