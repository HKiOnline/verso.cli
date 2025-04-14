package verext

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestParse(t *testing.T) {

	testDataDir := "../../tests/data"
	testFile := "keep_a_changelog.md"
	filePath := filepath.Join(testDataDir, testFile)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Fatalf("Test file not found: %s", filePath)
	}

	cl, err := Parse(filePath)

	if err != nil {
		t.Error(err)
	}

	if len(cl.Versions) == 0 {
		t.Error("No version extracted")
	}

	log.Printf("%d versions found", len(cl.Versions))
	log.Printf("latest version: %v+", cl.Versions[0])

	log.Printf("all versions:\n%v+", cl.Versions)
}
