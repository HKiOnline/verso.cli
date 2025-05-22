package output

import (
	"testing"
	"time"

	"github.com/hkionline/verso"
)

func TestLatest(t *testing.T) {
	t.Run("empty changelog", func(t *testing.T) {
		changelog := &verso.Changelog{}
		expected := ""
		actual := Latest(changelog)
		if actual != expected {
			t.Errorf("Expected: %q, Actual: %q", expected, actual)
		}
	})

	t.Run("single version", func(t *testing.T) {
		changelog := &verso.Changelog{
			Versions: []verso.Semver{
				{Major: 1, Minor: 2, Patch: 3},
			},
		}
		expected := "1.2.3\n"
		actual := Latest(changelog)
		if actual != expected {
			t.Errorf("Expected: %q, Actual: %q", expected, actual)
		}
	})

	t.Run("multiple versions", func(t *testing.T) {
		changelog := &verso.Changelog{
			Versions: []verso.Semver{
				{Major: 1, Minor: 2, Patch: 3},
				{Major: 4, Minor: 5, Patch: 6},
				{Major: 7, Minor: 8, Patch: 9},
			},
		}
		expected := "1.2.3\n"
		actual := Latest(changelog)
		if actual != expected {
			t.Errorf("Expected: %q, Actual: %q", expected, actual)
		}
	})

	t.Run("version with pre-release and build metadata", func(t *testing.T) {
		now := time.Now()
		changelog := &verso.Changelog{
			Versions: []verso.Semver{
				{Major: 1, Minor: 0, Patch: 0, PreRelease: "alpha", Build: "123", Date: now},
			},
		}
		expected := "1.0.0\n"
		actual := Latest(changelog)
		if actual != expected {
			t.Errorf("Expected: %q, Actual: %q", expected, actual)
		}
	})
}
