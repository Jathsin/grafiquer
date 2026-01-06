package web

import (
	"jathsin/utils"
	"net/http"

	"github.com/a-h/templ"
)

func Get_mux() (*http.ServeMux, error) {

	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("GET /about", about_handler)

	return mux, nil
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	if utils.IsHTMX(r) {
		templ.Handler(content_about()).ServeHTTP(w, r)
		return
	}
	templ.Handler(layout(nil, nav_bar(), content_about())).ServeHTTP(w, r)
}
