package book_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/austien/logbook/book"
)

func TestReadDay(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	homeDir := filepath.Join(cwd, "..", "testdata")

	dayFile, err := os.Open(filepath.Join(homeDir, "2024", "07", "24.md"))
	if err != nil {
		t.Fatal(err)
	}
	defer dayFile.Close()

	day, err := book.ReadDay(dayFile)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("got day %+v\n", day)
}
