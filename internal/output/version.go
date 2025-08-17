package output

import (
	"fmt"

	"github.com/hkionline/verso"
)

func Version() string {
	version, _ := getVersion()
	return fmt.Sprintf("v%d.%d.%d", version.Major, version.Minor, version.Patch)
}

func getVersion() (verso.Semver, error) {
	return verso.Semver{
		Major: 0,
		Minor: 5,
		Patch: 0,
	}, nil
}
