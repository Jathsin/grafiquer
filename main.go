package main

import (
	"context"
	"jathsin/landing"
	"jathsin/web"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

var log *slog.Logger

func main() {

	log = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	landing_mux, err := landing.Get_mux()
	if err != nil {
		log.Error("main: error in landing.Get_mux()", "error", err)
		os.Exit(1)
	}

	web_mux, err := web.Get_mux()
	if err != nil {
		log.Error("main: error in web.Get_mux()", "error", err)
		os.Exit(1)
	}

	mux := http.NewServeMux()

	mux.Handle("GET /{$}", landing_mux)

	mux.Handle("GET /{path...}", web_mux)

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
		log.Error("Error in server.ListenAndServe()", "error", err)
	}
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
