package main

import (
	"log"

	"github.com/hkionline/verso/internal/app"
)

func main() {

	app, err := app.New()
	if err != nil {
		log.Fatalf("failed to initialize the app: %v\n", err)
	}

	app.ReadChangelog()
	app.Output()

}
