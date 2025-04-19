package app

import (
	"flag"
	"log"
	"os"

	"github.com/hkionline/verso"
)

const defaultArg string = "latest"

type app struct {
	arg       string
	changelog verso.Changelog
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

func (a *app) ReadChangelog() {
	// TODO: Add logic to read CHANGELOG from a specific path. For now, current working directory is used.

	execPath, err := os.Executable()
	if err != nil {
		log.Fatalf("failed to get path for current executable: %v\n", err)
	}

	a.changelog, err = verso.Parse(execPath + "/CHANGELOG.md")

	if err != nil {
		log.Fatalf("failed to read version info: %v\n", err)
	}
}

func (a *app) Output() {
	panic("unimplemented")
}

func getFlagArg() string {
	flag.Parse()

	// TODO: Specify CHANGELOG path as an argument.

	args := flag.Args()

	if len(args) == 0 {
		return defaultArg
	}

	return args[0]
}
