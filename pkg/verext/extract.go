package verext

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

func extract(content string) (Changelog, error) {

	changelog := Changelog{
		Versions: []Semver{},
	}

	// Regular expression to match version headers
	re := regexp.MustCompile(`## \[(\d+\.\d+\.\d+(?:-\w+)?)\] - (\d{4}-\d{2}-\d{2})`)
	matches := re.FindAllStringSubmatch(content, -1)

	for _, match := range matches {

		if len(match) != 3 {
			continue // Skip invalid matches
		}

		versionStr := match[1]
		dateStr := match[2]

		// Parse version string
		parts := strings.Split(versionStr, ".")
		if len(parts) < 3 {
			continue // Skip invalid version strings
		}

		major, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}
		minor, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		patch, err := strconv.Atoi(parts[2])
		if err != nil {
			continue
		}

		// Handle pre-release versions (e.g., 1.2.3-alpha.1)
		preRelease := ""
		if strings.Contains(versionStr, "-") {
			preRelease = strings.Split(versionStr, "-")[1]
		}

		// Parse date
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			continue
		}

		// Create Semver struct
		semver := Semver{
			Major:      major,
			Minor:      minor,
			Patch:      patch,
			PreRelease: preRelease,
			Build:      "", // Not present in the example format
			Date:       date,
		}

		changelog.Versions = append(changelog.Versions, semver)
	}

	return changelog, nil
}
