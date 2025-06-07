package output

import (
	"testing"

	"github.com/hkionline/verso"
)

func TestBump(t *testing.T) {
	t.Run("empty changelog", func(t *testing.T) {
		changelog := &verso.Changelog{}
		expected := ""
		actual := Bump(changelog, Patch)
		if actual != expected {
			t.Errorf("Expected: %q, Actual: %q", expected, actual)
		}
	})

	t.Run("bump patch", func(t *testing.T) {
		changelog := &verso.Changelog{
			Versions: []verso.Semver{
				{Major: 1, Minor: 2, Patch: 3},
			},
		}
		expected := "1.2.4\n"
		actual := Bump(changelog, Patch)
		if actual != expected {
			t.Errorf("Expected: %q, Actual: %q", expected, actual)
		}
	})

	t.Run("bump minor", func(t *testing.T) {
		changelog := &verso.Changelog{
			Versions: []verso.Semver{
				{Major: 1, Minor: 2, Patch: 3},
			},
		}
		expected := "1.3.0\n"
		actual := Bump(changelog, Minor)
		if actual != expected {
			t.Errorf("Expected: %q, Actual: %q", expected, actual)
		}
	})

	t.Run("bump major", func(t *testing.T) {
		changelog := &verso.Changelog{
			Versions: []verso.Semver{
				{Major: 1, Minor: 2, Patch: 3},
			},
		}
		expected := "2.0.0\n"
		actual := Bump(changelog, Major)
		if actual != expected {
			t.Errorf("Expected: %q, Actual: %q", expected, actual)
		}
	})
}
