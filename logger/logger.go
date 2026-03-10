package logger

import (
	"context"
	"jathsin/types"
	"log/slog"
	"net/http"
	"os"

	"github.com/google/uuid"
)

var Log *slog.Logger

func Init() {
	Log = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

// TODO: rename parameter
func Logging(mux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		requestLogger := Log.With("request_id", id)
		requestLogger.Info("REQUEST",
			"method", r.Method,
			"url", r.URL.Path,
			"address", r.RemoteAddr,
		)

		// Attatch log data to request to later use the same log id across
		// packages for the same request
		ctx := context.WithValue(r.Context(), types.Ctx_key_logger{}, requestLogger)
		r = r.WithContext(ctx)

		mux.ServeHTTP(w, r)
	})
}
