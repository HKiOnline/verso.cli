package output

import "github.com/hkionline/verso"

func Version() string {
	version, _ := getVersion()
	return version.String()
}

func getVersion() (verso.Semver, error) {
	return verso.Semver{
		Major: 0,
		Minor: 5,
		Patch: 0,
	}, nil
}
