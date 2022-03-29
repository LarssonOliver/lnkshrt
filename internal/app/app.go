package app

import (
	"log"

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
