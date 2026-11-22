package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func get(t *testing.T, h http.Handler, path string) (int, string) {
	t.Helper()
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, path, nil))
	body, _ := io.ReadAll(rec.Body)
	return rec.Code, string(body)
}

func TestVersionIsTheBuiltCommit(t *testing.T) {
	version = "abc123"
	if code, body := get(t, routes(t.TempDir()), "/version"); code != 200 || body != "abc123" {
		t.Fatalf("got %d %q", code, body)
	}
}

func TestHealth(t *testing.T) {
	if code, body := get(t, routes(t.TempDir()), "/api/health"); code != 200 || body != `{"ok":true}` {
		t.Fatalf("got %d %q", code, body)
	}
}

func TestServesTheWebApp(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "index.html"), []byte("<h1>meal-prep</h1>"), 0o644); err != nil {
		t.Fatal(err)
	}
	if code, body := get(t, routes(dir), "/"); code != 200 || body != "<h1>meal-prep</h1>" {
		t.Fatalf("got %d %q", code, body)
	}
}
