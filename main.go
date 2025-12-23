package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/a-h/templ"
	"github.com/google/uuid"
)

var log *slog.Logger

func main() {

	log = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	mux := http.NewServeMux()

	// Hypermedia API

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.HandleFunc("GET /", landing_handler)

	// Build server
	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      logging(mux),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Error("Error in server.ListenAndServe()", "error", err)
	}
}

// ---------------------------- Hypermedia API -----------------------------

func landing_handler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(layout(perlin())).ServeHTTP(w, r)
}

func logging(f http.Handler) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		id := uuid.New().String()
		request_log := log.With("request_id", id)
		request_log.Info("REQUEST",
			"method", r.Method,
			"url", r.URL.Path,
			"address", r.RemoteAddr)

		ctx := context.WithValue(r.Context(), "logs", request_log)
		r = r.WithContext(ctx)

		f.ServeHTTP(w, r)
	})
}
