package output

import (
	"fmt"
	"github.com/hkionline/verso"
)

// Output the latest version found from the changelog
func Latest(changelog *verso.Changelog) string {
	version := changelog.Versions[0]
	return fmt.Sprintf("%d.%d.%d\n", version.Major, version.Minor, version.Patch)
}
