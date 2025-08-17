package output

import "strings"

func Help() string {

	sb := strings.Builder{}

	sb.WriteString("\n")
	sb.WriteString("Verso CLI - Simple semantic version extractor CLI ")
	sb.WriteString(Version())
	sb.WriteString("\n\n")

	sb.WriteString("Usage:\n")
	sb.WriteString("\n")
	sb.WriteString("\tverso latest - get the latest version from the changelog\n")
	sb.WriteString("\tverso list - get a list of version found from a changelog\n")
	sb.WriteString("\tverso bump major - get a new bumped major version\n")
	sb.WriteString("\tverso bump minor - get a new bumped minor version\n")
	sb.WriteString("\tverso bump patch - get a new bumped patch version\n")

	return sb.String()
}
