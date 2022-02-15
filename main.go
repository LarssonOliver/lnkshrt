package main

import (
	"log"
	"net/http"

	"larssonoliver.com/lnkshrt/app"
	"larssonoliver.com/lnkshrt/app/config"
)

const port = ":8080"

func logInit() {
	log.SetPrefix("\033[35m[" + config.Executable() + "]:\033[0m ")
}

func main() {
	logInit()
	app := app.New()

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, app.Router))
}
