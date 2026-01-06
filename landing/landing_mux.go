package landing

import (
	"net/http"

	"jathsin/utils"

	"github.com/a-h/templ"
)

func Get_mux() (*http.ServeMux, error) {

	mux := http.NewServeMux()

	// Redirect to landing by default
	mux.HandleFunc("GET /", landing_handler)

	return mux, nil
}

func landing_handler(w http.ResponseWriter, r *http.Request) {
	if utils.IsHTMX(r) {
		templ.Handler(perlin()).ServeHTTP(w, r)
		return
	}
	templ.Handler(layout(perlin(), nav_bar(), nil)).ServeHTTP(w, r)
}
