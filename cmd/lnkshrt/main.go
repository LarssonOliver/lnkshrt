package main

import (
	"log"

	"github.com/larssonoliver/lnkshrt/internal/app"
	"github.com/larssonoliver/lnkshrt/internal/config"
)

func logInit() {
	log.SetPrefix("\033[35m[" + config.Executable() + "]:\033[0m ")
}

func main() {
	logInit()
	app.New().Run()
}
