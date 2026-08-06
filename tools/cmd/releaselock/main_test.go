package main

import (
	"archive/zip"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestVerifyEntry(t *testing.T) {
	t.Parallel()
	path := writeTestArchive(t, []string{archiveEntry})
	if err := verifyEntry(path); err != nil {
		t.Fatal(err)
	}
}

func TestVerifyEntryRejectsDuplicate(t *testing.T) {
	t.Parallel()
	path := writeTestArchive(t, []string{archiveEntry, archiveEntry})
	if err := verifyEntry(path); err == nil || !strings.Contains(err.Error(), "duplicate") {
		t.Fatalf("expected duplicate entry error, got %v", err)
	}
}

func TestVerifyEntryRejectsMissing(t *testing.T) {
	t.Parallel()
	path := writeTestArchive(t, []string{"not-the-spec.json"})
	if err := verifyEntry(path); err == nil || !strings.Contains(err.Error(), "not found") {
		t.Fatalf("expected missing entry error, got %v", err)
	}
}

func writeTestArchive(t *testing.T, entries []string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "artifact.war")
	file, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	archive := zip.NewWriter(file)
	for _, entry := range entries {
		writer, err := archive.Create(entry)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := writer.Write([]byte(`{"openapi":"3.0.1"}`)); err != nil {
			t.Fatal(err)
		}
	}
	if err := archive.Close(); err != nil {
		t.Fatal(err)
	}
	if err := file.Close(); err != nil {
		t.Fatal(err)
	}
	return path
}
