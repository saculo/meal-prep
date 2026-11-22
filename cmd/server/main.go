// The meal-prep server: the API and, from STATIC_DIR, the built web app (web/dist).
package main

import (
	"log"
	"net/http"
	"os"
)

// version is the commit the binary was built from (-ldflags "-X main.version=<sha>").
var version = "dev"

func routes(static string) http.Handler {
	mux := http.NewServeMux()
	// what is deployed: a preview is ready for a commit when this returns its SHA
	mux.HandleFunc("GET /version", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = w.Write([]byte(version))
	})
	mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true}`))
	})
	mux.Handle("/", http.FileServer(http.Dir(static)))
	return mux
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func main() {
	addr := ":" + env("PORT", "8080")
	log.Printf("meal-prep %s on %s", version, addr)
	log.Fatal(http.ListenAndServe(addr, routes(env("STATIC_DIR", "web/dist"))))
}
