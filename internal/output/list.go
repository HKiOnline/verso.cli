package output

import (
	"fmt"
	"strings"

	"github.com/hkionline/verso"
)

// Output a list of all versions found from the changelog.
// Each version will have its own line.
func List(changelog *verso.Changelog) string {

	sb := strings.Builder{}

	for _, version := range changelog.Versions {
		sb.WriteString(fmt.Sprintf("%d.%d.%d\n", version.Major, version.Minor, version.Patch))
	}

	return sb.String()
}
