package app

import (
	"github.com/gorilla/mux"
	"larssonoliver.com/lnkshrt/app/db"
)

type App struct {
	Router   *mux.Router
	Database *db.Database
}

func New() *App {
	a := &App{
		Router:   mux.NewRouter().StrictSlash(true),
		Database: db.New(),
	}

	a.initRoutes()
	return a
}
