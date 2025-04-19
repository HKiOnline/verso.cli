package main

import (
	"fmt"
	"log"

	"github.com/hkionline/verso/internal/app"
)

func main() {

	app, err := app.New()
	if err != nil {
		log.Fatalf("failed to initialize the app: %v\n", err)
	}

	fmt.Printf("app: %+v\n", app)

}
