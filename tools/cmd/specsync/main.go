// Command specsync downloads, verifies, extracts, normalizes, and validates
// the OpenAPI documents published in Apache NiFi release WARs.
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
	"sort"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
)

const maxSpecBytes = 32 << 20

var versionPattern = regexp.MustCompile(`^2\.[0-9]+\.[0-9]+$`)

type sourceLock struct {
	Version   string              `json:"version"`
	Artifacts map[string]artifact `json:"artifacts"`
}

type artifact struct {
	URL      string `json:"url"`
	Entry    string `json:"entry"`
	SHA512   string `json:"sha512"`
	MaxBytes int64  `json:"maxBytes"`
}

func main() {
	var lockPath string
	var cacheDir string
	var outputDir string
	flag.StringVar(&lockPath, "lock", "source.lock.json", "path to the source lock")
	flag.StringVar(&cacheDir, "cache", ".cache/artifacts", "artifact cache directory")
	flag.StringVar(&outputDir, "output", "openapi", "OpenAPI output directory")
	flag.Parse()

	if err := run(lockPath, cacheDir, outputDir); err != nil {
		fmt.Fprintln(os.Stderr, "specsync:", err)
		os.Exit(1)
	}
}

func run(lockPath, cacheDir, outputDir string) error {
	lockData, err := os.ReadFile(lockPath)
	if err != nil {
		return fmt.Errorf("read lock: %w", err)
	}

	var lock sourceLock
	if err := json.Unmarshal(lockData, &lock); err != nil {
		return fmt.Errorf("decode lock: %w", err)
	}
	if !versionPattern.MatchString(lock.Version) {
		return fmt.Errorf("unsupported version %q", lock.Version)
	}
	if len(lock.Artifacts) != 2 {
		return fmt.Errorf("lock must contain exactly nifi and registry artifacts")
	}

	for _, name := range []string{"nifi", "registry"} {
		a, ok := lock.Artifacts[name]
		if !ok {
			return fmt.Errorf("lock is missing %s artifact", name)
		}
		if err := syncArtifact(name, lock.Version, a, cacheDir, outputDir); err != nil {
			return err
		}
	}
	return nil
}

func syncArtifact(name, version string, a artifact, cacheDir, outputDir string) error {
	if err := validateArtifactSource(name, version, a.URL); err != nil {
		return err
	}
	if a.MaxBytes <= 0 || a.MaxBytes > 256<<20 {
		return fmt.Errorf("%s maxBytes must be between 1 and 256 MiB", name)
	}
	if len(a.SHA512) != sha512.Size*2 {
		return fmt.Errorf("%s has invalid SHA-512 length", name)
	}
	if _, err := hex.DecodeString(a.SHA512); err != nil {
		return fmt.Errorf("%s has invalid SHA-512: %w", name, err)
	}
	if a.Entry == "" || strings.HasPrefix(a.Entry, "/") || strings.Contains(a.Entry, "..") {
		return fmt.Errorf("%s has unsafe archive entry %q", name, a.Entry)
	}

	if err := os.MkdirAll(cacheDir, 0o755); err != nil {
		return fmt.Errorf("create artifact cache: %w", err)
	}
	archivePath := filepath.Join(cacheDir, a.SHA512+".war")
	if err := ensureArtifact(archivePath, a); err != nil {
		return fmt.Errorf("%s artifact: %w", name, err)
	}

	raw, err := extract(archivePath, a.Entry)
	if err != nil {
		return fmt.Errorf("%s extract: %w", name, err)
	}
	normalized, err := normalize(raw, version)
	if err != nil {
		return fmt.Errorf("%s normalize: %w", name, err)
	}
	if err := validate(normalized, version); err != nil {
		return fmt.Errorf("%s validate: %w", name, err)
	}

	if err := writeAtomic(filepath.Join(outputDir, "raw", name+".json"), raw); err != nil {
		return err
	}
	if err := writeAtomic(filepath.Join(outputDir, "normalized", name+".json"), normalized); err != nil {
		return err
	}
	fmt.Printf("synced %s OpenAPI %s\n", name, version)
	return nil
}

func validateArtifactSource(name, version, rawURL string) error {
	var expected string
	switch name {
	case "nifi":
		expected = fmt.Sprintf("https://repo1.maven.org/maven2/org/apache/nifi/nifi-web-api/%[1]s/nifi-web-api-%[1]s.war", version)
	case "registry":
		expected = fmt.Sprintf("https://repo1.maven.org/maven2/org/apache/nifi/registry/nifi-registry-web-api/%[1]s/nifi-registry-web-api-%[1]s.war", version)
	default:
		return fmt.Errorf("unknown artifact %q", name)
	}
	if rawURL != expected {
		return fmt.Errorf("%s URL must be the exact Maven Central release artifact: got %q", name, rawURL)
	}
	return nil
}

func ensureArtifact(path string, a artifact) error {
	if f, err := os.Open(path); err == nil {
		defer f.Close()
		if digest, size, err := hashReader(f, a.MaxBytes); err == nil && digest == a.SHA512 && size <= a.MaxBytes {
			return nil
		}
	}

	client := &http.Client{
		Timeout: 5 * time.Minute,
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			return errors.New("artifact download redirects are not allowed")
		},
	}
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, a.URL, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download returned %s", resp.Status)
	}

	tmp, err := os.CreateTemp(filepath.Dir(path), ".download-*")
	if err != nil {
		return err
	}
	tmpPath := tmp.Name()
	defer os.Remove(tmpPath)

	h := sha512.New()
	written, copyErr := io.Copy(io.MultiWriter(tmp, h), io.LimitReader(resp.Body, a.MaxBytes+1))
	closeErr := tmp.Close()
	if copyErr != nil {
		return copyErr
	}
	if closeErr != nil {
		return closeErr
	}
	if written > a.MaxBytes {
		return fmt.Errorf("download exceeds %d bytes", a.MaxBytes)
	}
	if got := hex.EncodeToString(h.Sum(nil)); got != a.SHA512 {
		return fmt.Errorf("SHA-512 mismatch: got %s", got)
	}
	if err := os.Rename(tmpPath, path); err != nil {
		return err
	}
	return nil
}

func hashReader(r io.Reader, limit int64) (string, int64, error) {
	h := sha512.New()
	n, err := io.Copy(h, io.LimitReader(r, limit+1))
	if err != nil {
		return "", n, err
	}
	return hex.EncodeToString(h.Sum(nil)), n, nil
}

func extract(archivePath, entry string) ([]byte, error) {
	zr, err := zip.OpenReader(archivePath)
	if err != nil {
		return nil, err
	}
	defer zr.Close()
	var selected *zip.File
	for _, f := range zr.File {
		if f.Name != entry {
			continue
		}
		if selected != nil {
			return nil, fmt.Errorf("archive contains duplicate entry %q", entry)
		}
		selected = f
	}
	if selected == nil {
		return nil, fmt.Errorf("entry %q not found", entry)
	}
	if selected.UncompressedSize64 > maxSpecBytes {
		return nil, fmt.Errorf("spec exceeds %d bytes", maxSpecBytes)
	}
	r, err := selected.Open()
	if err != nil {
		return nil, err
	}
	defer r.Close()
	data, err := io.ReadAll(io.LimitReader(r, maxSpecBytes+1))
	if err != nil {
		return nil, err
	}
	if len(data) > maxSpecBytes {
		return nil, fmt.Errorf("spec exceeds %d bytes", maxSpecBytes)
	}
	return data, nil
}

func normalize(raw []byte, expectedVersion string) ([]byte, error) {
	var doc map[string]any
	if err := json.Unmarshal(raw, &doc); err != nil {
		return nil, err
	}
	if versionOf(doc) != expectedVersion {
		return nil, fmt.Errorf("spec version %q does not match %q", versionOf(doc), expectedVersion)
	}

	declared := map[string]struct{}{}
	if components, ok := object(doc["components"]); ok {
		if schemes, ok := object(components["securitySchemes"]); ok {
			for name, value := range schemes {
				declared[name] = struct{}{}
				scheme, ok := object(value)
				if !ok || scheme["type"] != "http" {
					continue
				}
				if s, ok := scheme["scheme"].(string); ok {
					scheme["scheme"] = strings.ToLower(s)
				}
			}
		}
	}

	paths, ok := object(doc["paths"])
	if !ok {
		return nil, errors.New("missing paths object")
	}
	operationIDs := map[string]string{}
	for pathName, pathValue := range paths {
		pathItem, ok := object(pathValue)
		if !ok {
			continue
		}
		for method, operationValue := range pathItem {
			if !isHTTPMethod(method) {
				continue
			}
			operation, ok := object(operationValue)
			if !ok {
				return nil, fmt.Errorf("%s %s is not an object", strings.ToUpper(method), pathName)
			}
			operationID, _ := operation["operationId"].(string)
			if operationID == "" {
				return nil, fmt.Errorf("%s %s has no operationId", strings.ToUpper(method), pathName)
			}
			if previous, exists := operationIDs[operationID]; exists {
				return nil, fmt.Errorf("duplicate operationId %q at %s and %s", operationID, previous, pathName)
			}
			operationIDs[operationID] = pathName
			normalizeResponses(operation)
			normalizeSecurity(operation, declared)
		}
	}

	out, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(out, '\n'), nil
}

func normalizeResponses(operation map[string]any) {
	responses, ok := object(operation["responses"])
	if !ok {
		return
	}
	for status, responseValue := range responses {
		response, ok := object(responseValue)
		if !ok {
			continue
		}
		if description, _ := response["description"].(string); description != "" {
			continue
		}
		if status == "default" {
			response["description"] = "Default response"
		} else {
			response["description"] = "HTTP " + status + " response"
		}
	}
}

func normalizeSecurity(operation map[string]any, declared map[string]struct{}) {
	security, present := operation["security"].([]any)
	if !present || len(security) == 0 {
		return
	}
	unknown := map[string]struct{}{}
	valid := make([]any, 0, len(security))
	for _, requirementValue := range security {
		requirement, ok := object(requirementValue)
		if !ok {
			continue
		}
		kept := map[string]any{}
		for name, scopes := range requirement {
			if _, ok := declared[name]; ok {
				kept[name] = scopes
			} else {
				unknown[name] = struct{}{}
			}
		}
		if len(kept) > 0 {
			valid = append(valid, kept)
		}
	}
	if len(valid) == 0 {
		delete(operation, "security")
	} else {
		operation["security"] = valid
	}
	if len(unknown) > 0 {
		authorizations := make([]string, 0, len(unknown))
		for name := range unknown {
			authorizations = append(authorizations, name)
		}
		sort.Strings(authorizations)
		operation["x-nifi-authorizations"] = authorizations
	}
}

func validate(data []byte, expectedVersion string) error {
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromData(data)
	if err != nil {
		return err
	}
	if doc.Info == nil || doc.Info.Version != expectedVersion {
		return fmt.Errorf("validated spec version does not match %q", expectedVersion)
	}
	return doc.Validate(context.Background())
}

func versionOf(doc map[string]any) string {
	info, _ := object(doc["info"])
	version, _ := info["version"].(string)
	return version
}

func object(value any) (map[string]any, bool) {
	result, ok := value.(map[string]any)
	return result, ok
}

func isHTTPMethod(method string) bool {
	switch method {
	case "get", "put", "post", "delete", "options", "head", "patch", "trace":
		return true
	default:
		return false
	}
}

func writeAtomic(path string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	tmp, err := os.CreateTemp(filepath.Dir(path), ".spec-*")
	if err != nil {
		return err
	}
	tmpPath := tmp.Name()
	defer os.Remove(tmpPath)
	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	if err := os.Chmod(tmpPath, 0o644); err != nil {
		return err
	}
	return os.Rename(tmpPath, path)
}
