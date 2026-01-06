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

	mux.HandleFunc("GET /articles", still_working)

	mux.HandleFunc("GET /projects", still_working)

	return mux, nil
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	if utils.IsHTMX(r) {
		templ.Handler(about()).ServeHTTP(w, r)
		return
	}
	templ.Handler(layout(nil, nav_bar(), about())).ServeHTTP(w, r)
}

// func projects_handler(w http.ResponseWriter, r *http.Request) {
// 	if utils.IsHTMX(r) {
// 		templ.Handler(about()).ServeHTTP(w, r)
// 		return
// 	}
// 	templ.Handler(layout(nil, nav_bar(), about())).ServeHTTP(w, r)
// }

func still_working(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	if utils.IsHTMX(r) {
		templ.Handler(not_found()).ServeHTTP(w, r)
		return
	}
	templ.Handler(layout(nil, nav_bar(), not_found())).ServeHTTP(w, r)
}
