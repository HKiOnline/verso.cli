package app

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/hkionline/verso"
	"github.com/hkionline/verso.cli/internal/output"
)

type app struct {
	args      []string
	changelog verso.Changelog
}

// New initializes and returns the app struct. Possible command-line arguments and flags
// are parsed, and if not present, default output is produced.
func New() (*app, error) {

	app := &app{
		args: getFlagArg(),
	}

	// TODO: Handle errors - if logic doesn't need to handle errors, remove the error return value!

	return app, nil
}

// ReadChangelog reads a changelog either from stdin if there is content or from working directory
func (a *app) ReadChangelog() {

	stat, err := os.Stdin.Stat()

	// Highly unlikely to happen, but let's deal with it anyways
	if err != nil {
		log.Fatalf("failed to get stats from stdin: %s", err)
	}

	// if we have something in stdin then we'll read the changelog from there
	// otherwise seek a changelog from the working directory
	if stat.Size() > 0 {
		a.readFromStdIn()
	} else {
		a.readFromWorkingDirectory()
	}

}

// ReadFromWorkingDirectory reads the CHANGELOG.md file from the current path.
func (a *app) readFromWorkingDirectory() {
	// TODO: Add logic to read CHANGELOG from a specific path. For now, current working directory is used.

	execPath, err := os.Getwd()

	if err != nil {
		log.Fatalf("failed to get path for current working directory: %v\n", err)
	}

	a.changelog, err = verso.ParsePath(execPath + "/CHANGELOG.md")

	if err != nil {
		log.Fatalf("failed to read version info: %v\n", err)
	}
}

// ReadFromStdIn reads contents of a changelog from the standard in
func (a *app) readFromStdIn() {

	bytes, err := io.ReadAll(os.Stdin)

	if err != nil {
		log.Fatalf("error reading changelog from stdin: %s", err)
	}

	a.changelog, err = verso.ParseBytes(bytes)

	if err != nil {
		log.Fatalf("error parsing changelog from bytes: %s", err)
	}
}

func (a *app) Output() {

	if len(a.args) < 1 {
		a.args = []string{"invalid"}
	}

	mainCmd := a.args[0]

	switch mainCmd {

	case "latest":
		fallthrough
	case "l":
		fmt.Fprint(os.Stdout, output.Latest(&a.changelog))
	case "list":
		fallthrough
	case "ls":
		fmt.Fprint(os.Stdout, output.List(&a.changelog))
	case "bump":
		fallthrough
	case "b":

		if len(a.args) < 2 {
			a.args[1] = "invalid"
		}

		subCmd := a.args[1]

		switch subCmd {
		case "patch":
			fallthrough
		case "+":
			fmt.Fprint(os.Stdout, output.Bump(&a.changelog, output.Patch))
		case "minor":
			fallthrough
		case "++":
			fmt.Fprint(os.Stdout, output.Bump(&a.changelog, output.Minor))
		case "major":
			fallthrough
		case "+++":
			fmt.Fprint(os.Stdout, output.Bump(&a.changelog, output.Major))
		default:
			fmt.Fprint(os.Stderr, "bump command requires a sub-command: use either bump patch, bump minor or bump major")
		}
	default:
		fmt.Fprint(os.Stdout, output.Latest(&a.changelog))
	}

}

func getFlagArg() []string {
	flag.Parse()
	return flag.Args()
}
