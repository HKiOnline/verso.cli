package main

import (
	"github.com/hkionline/verso.cli/internal/app"
	"log"
)

func main() {
	app, err := app.New()

	if err != nil {
		log.Fatalf("failed to initialize the app: %v\n", err)
	}

	app.ReadChangelog()
	app.Output()
}
