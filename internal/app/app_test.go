package app

import (
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("basic initialization", func(t *testing.T) {
		app, err := New()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if app == nil {
			t.Errorf("App should not be nil")
		}

		// Check that the args slice is initialized (even if it's empty)
		if app.args == nil {
			t.Errorf("Args slice should not be nil")
		}

		// Check that the changelog is initialized (even if it's empty)
		if app.changelog.Versions == nil {
			t.Errorf("Changelog versions slice should not be nil")
		}
	})
}

