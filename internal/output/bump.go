package output

import (
	"fmt"

	"github.com/hkionline/verso"
)

// Output a incremented version (bumped) of the latest version
// found from the changelog.
func Bump(changelog *verso.Changelog, incMajor bool, incMinor bool, incPatch bool) string {

	version := changelog.Versions[0]

	if incMajor {
		version.Major += 1
	}
	if incMinor {
		version.Minor += 1
	}
	if incPatch {
		version.Patch += 1
	}

	return fmt.Sprintf("%d.%d.%d\n", version.Major, version.Minor, version.Patch)

}
