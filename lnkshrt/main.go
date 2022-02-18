package main

import (
	"fmt"
	"log"
	"net/http"

	"larssonoliver.com/lnkshrt/app"
	"larssonoliver.com/lnkshrt/app/config"
)

func logInit() {
	log.SetPrefix("\033[35m[" + config.Executable() + "]:\033[0m ")
}

func main() {
	logInit()
	app := app.New()

	port := fmt.Sprintf(":%d", config.Port())

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, app.Router))
}
