// Command releaselock creates a source lock for an exact Apache NiFi release.
package main

import (
	"archive/zip"
	"context"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

const archiveEntry = "docs/rest-api/swagger.json"
const userAgent = "nifi-go-bindings/2 (+https://github.com/skycubed/nifi-go)"

var versionPattern = regexp.MustCompile(`^2\.[0-9]+\.[0-9]+$`)

type artifact struct {
	URL      string `json:"url"`
	Entry    string `json:"entry"`
	SHA512   string `json:"sha512"`
	MaxBytes int64  `json:"maxBytes"`
}

func main() {
	var version string
	var output string
	flag.StringVar(&version, "version", "", "Apache NiFi release version")
	flag.StringVar(&output, "output", "source.lock.json", "output lock path")
	flag.Parse()
	if !versionPattern.MatchString(version) {
		fail(fmt.Errorf("invalid version %q", version))
	}

	artifacts := map[string]artifact{
		"nifi": {
			URL:      fmt.Sprintf("https://repo.maven.apache.org/maven2/org/apache/nifi/nifi-web-api/%[1]s/nifi-web-api-%[1]s.war", version),
			Entry:    archiveEntry,
			MaxBytes: 10 << 20,
		},
		"registry": {
			URL:      fmt.Sprintf("https://repo.maven.apache.org/maven2/org/apache/nifi/registry/nifi-registry-web-api/%[1]s/nifi-registry-web-api-%[1]s.war", version),
			Entry:    archiveEntry,
			MaxBytes: 100 << 20,
		},
	}

	temporaryDirectory, err := os.MkdirTemp("", "nifi-release-lock-*")
	if err != nil {
		fail(err)
	}
	defer os.RemoveAll(temporaryDirectory)

	for name, value := range artifacts {
		digest, err := downloadAndVerify(value.URL, value.MaxBytes, filepath.Join(temporaryDirectory, name+".war"))
		if err != nil {
			fail(fmt.Errorf("%s: %w", name, err))
		}
		value.SHA512 = digest
		artifacts[name] = value
	}

	lock := struct {
		Version   string              `json:"version"`
		Artifacts map[string]artifact `json:"artifacts"`
	}{Version: version, Artifacts: artifacts}
	data, err := json.MarshalIndent(lock, "", "  ")
	if err != nil {
		fail(err)
	}
	data = append(data, '\n')
	if err := writeAtomic(output, data); err != nil {
		fail(err)
	}
}

func writeAtomic(path string, data []byte) error {
	directory := filepath.Dir(path)
	temporary, err := os.CreateTemp(directory, ".source-lock-*")
	if err != nil {
		return err
	}
	temporaryPath := temporary.Name()
	defer os.Remove(temporaryPath)
	if _, err := temporary.Write(data); err != nil {
		temporary.Close()
		return err
	}
	if err := temporary.Close(); err != nil {
		return err
	}
	if err := os.Chmod(temporaryPath, 0o644); err != nil {
		return err
	}
	return os.Rename(temporaryPath, path)
}

func downloadAndVerify(url string, maximum int64, output string) (string, error) {
	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	request.Header.Set("User-Agent", userAgent)
	client := &http.Client{
		Timeout: 5 * time.Minute,
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			return errors.New("artifact download redirects are not allowed")
		},
	}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("download returned %s", response.Status)
	}

	file, err := os.Create(output)
	if err != nil {
		return "", err
	}
	hash := sha512.New()
	written, copyErr := io.Copy(io.MultiWriter(file, hash), io.LimitReader(response.Body, maximum+1))
	closeErr := file.Close()
	if copyErr != nil {
		return "", copyErr
	}
	if closeErr != nil {
		return "", closeErr
	}
	if written > maximum {
		return "", fmt.Errorf("artifact exceeds %d bytes", maximum)
	}
	if err := verifyEntry(output); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func verifyEntry(path string) error {
	reader, err := zip.OpenReader(path)
	if err != nil {
		return err
	}
	defer reader.Close()
	found := false
	for _, file := range reader.File {
		if file.Name != archiveEntry {
			continue
		}
		if found {
			return fmt.Errorf("archive contains duplicate entry %q", archiveEntry)
		}
		if file.UncompressedSize64 > 32<<20 {
			return fmt.Errorf("archive entry %q exceeds 32 MiB", archiveEntry)
		}
		found = true
	}
	if !found {
		return fmt.Errorf("safe %q entry not found", archiveEntry)
	}
	return nil
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, "releaselock:", err)
	os.Exit(1)
}
