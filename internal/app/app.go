package app

import (
	"flag"
)

const defaultArg string = "latest"

type app struct {
	arg string
}

// New initializes and returns the app struct. Possible command-line arguments and flags
// are parsed, and if not present, default values are used.
func New() (*app, error) {

	app := &app{
		arg: getFlagArg(),
	}

	// TODO: Handle errors - if logic doesn't need to handle errors, remove the error return value!

	return app, nil
}

func getFlagArg() string {
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		return defaultArg
	}

	return args[0]
}
