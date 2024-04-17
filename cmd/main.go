package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"fmt"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
		"github.com/art3xias23/go-test/internal/components"
)

func main() {

	fmt.Print("Hello")
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	r.Get("/templ", templ.Handler(components.Hell("MyBoy")).ServeHTTP)

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "styles"))
	FileServer(r, "/", filesDir)

	http.ListenAndServe(":8080", r)

}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
