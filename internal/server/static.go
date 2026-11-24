package server

import (
	"io/fs"
	"net/http"
)

// spa serves the built UI, falling back to index.html for client routes.
func spa(static fs.FS) http.Handler {
	files := http.FileServerFS(static)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := fs.Stat(static, trimSlash(r.URL.Path)); err != nil {
			http.ServeFileFS(w, r, static, "index.html")
			return
		}
		files.ServeHTTP(w, r)
	})
}

func trimSlash(p string) string {
	if p == "/" || p == "" {
		return "index.html"
	}
	return p[1:]
}
