package output

import (
	"testing"
	"time"

	"github.com/hkionline/verso"
)

func TestList(t *testing.T) {
	t.Run("empty changelog", func(t *testing.T) {
		changelog := &verso.Changelog{}
		expected := ""
		actual := List(changelog)
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
		actual := List(changelog)
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
		expected := "1.2.3\n4.5.6\n7.8.9\n"
		actual := List(changelog)
		if actual != expected {
			t.Errorf("Expected: %q, Actual: %q", expected, actual)
		}
	})

	t.Run("versions with pre-release and build metadata", func(t *testing.T) {
		now := time.Now()
		changelog := &verso.Changelog{
			Versions: []verso.Semver{
				{Major: 1, Minor: 0, Patch: 0, PreRelease: "alpha", Build: "123", Date: now},
				{Major: 1, Minor: 0, Patch: 1, PreRelease: "beta", Build: "456", Date: now},
			},
		}
		expected := "1.0.0\n1.0.1\n" // The function only outputs major.minor.patch
		actual := List(changelog)
		if actual != expected {
			t.Errorf("Expected: %q, Actual: %q", expected, actual)
		}
	})
}
