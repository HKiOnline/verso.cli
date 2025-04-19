package app

import (
	"os"
	"testing"
	//"github.com/hkionline/verso/internal/app"
)

func TestNewApp(t *testing.T) {
	t.Parallel()

	firstArg := "test_app"

	tests := []struct {
		name   string
		args   []string
		expArg string
	}{
		{
			name:   "without args",
			args:   []string{firstArg},
			expArg: "latest",
		},
		{
			name:   "with arg 'latest'",
			args:   []string{firstArg, "latest"},
			expArg: "latest",
		},
		{
			name:   "with arg 'list'",
			args:   []string{firstArg, "list"},
			expArg: "list",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			os.Args = test.args

			app, err := New()

			if err != nil {
				t.Fatalf("creation of app shouldn't fail - error: %v\n", err)
			}

			if test.expArg != app.arg {
				t.Fatalf("expected flag argument \"%s\" but got \"%s\"\n", test.expArg, app.arg)
			}
		})
	}
}
