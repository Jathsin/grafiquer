package main

import (
	"context"
	"jathsin/landing"
	"jathsin/web/about"
	"jathsin/web/articles"
	"jathsin/web/projects"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

type ctx_key_logger struct{}

var logger *slog.Logger

func main() {

	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	landing_mux, err := landing.Get_mux()
	if err != nil {
		logger.Error("main: error in landing.Get_mux()", "error", err)
		os.Exit(1)
	}

	articles_mux, err := articles.Get_mux()
	if err != nil {
		logger.Error("main: error in articles.Get_mux()", "error", err)
		os.Exit(1)
	}

	projects_mux, err := projects.Get_mux()
	if err != nil {
		logger.Error("main: error in projects.Get_mux()", "error", err)
		os.Exit(1)
	}

	about_mux, err := about.Get_mux()
	if err != nil {
		logger.Error("main: error in web.Get_mux()", "error", err)
		os.Exit(1)
	}

	mux := http.NewServeMux()

	mux.Handle("GET /{$}", landing_mux)

	// Canonical URL for list page: /projects (no trailing slash)
	mux.Handle("GET /projects", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r2 := r.Clone(r.Context())
		r2.URL.Path = "/" // make projects_mux match its own "GET /"
		projects_mux.ServeHTTP(w, r2)
	}))

	// If someone hits /projects/ exactly, redirect to /projects, otherwise delegate to the subtree handler.
	mux.Handle("GET /projects/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/projects/" {
			http.Redirect(w, r, "/projects", http.StatusMovedPermanently)
			return
		}
		http.StripPrefix("/projects", projects_mux).ServeHTTP(w, r)
	}))

	mux.Handle("GET /about", about_mux)

	mux.Handle("GET /articles", articles_mux)

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Build server
	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      logging(mux),
	}

	err = server.ListenAndServe()
	if err != nil {
		logger.Error("Error in server.ListenAndServe()", "error", err)
	}
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		requestLogger := logger.With("request_id", id)
		requestLogger.Info("REQUEST",
			"method", r.Method,
			"url", r.URL.Path,
			"address", r.RemoteAddr,
		)

		// Attatch log data to request to later use the same log id across
		// packages for the same request
		ctx := context.WithValue(r.Context(), ctx_key_logger{}, requestLogger)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
