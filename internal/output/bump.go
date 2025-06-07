package output

import (
	"fmt"

	"github.com/hkionline/verso"
)

type BumpLevel int

const (
	Patch BumpLevel = iota + 1
	Minor
	Major
)

// Output a incremented version (bumped) of the latest version
// found from the changelog.
func Bump(changelog *verso.Changelog, bumpLevel BumpLevel) string {

	if len(changelog.Versions) < 1 {
		return ""
	}

	version := changelog.Versions[0]

	if bumpLevel == Patch {
		version.Patch += 1
	}

	if bumpLevel == Minor {
		version.Patch = 0
		version.Minor += 1
	}

	if bumpLevel == Major {
		version.Patch = 0
		version.Minor = 0
		version.Major += 1
	}

	return fmt.Sprintf("%d.%d.%d\n", version.Major, version.Minor, version.Patch)
}
