package projects

import (
	"jathsin/utils"
	"jathsin/web/ui"
	"net/http"
	"os"
	"path/filepath"

	"github.com/a-h/templ"
)

func Get_mux() (*http.ServeMux, error) {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", projects_handler)

	mux.HandleFunc("GET /{name}", show_project_handler)

	mux.HandleFunc("GET /{name}/static/{file...}", func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		file := r.PathValue("file")

		// validate name like you already do
		// then:
		http.ServeFile(w, r, filepath.Join("projects", name, "static", file))
	})

	mux.HandleFunc("GET /{name}/{file...}", func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		file := r.PathValue("file")

		// validate name like you already do
		// then:
		http.ServeFile(w, r, filepath.Join("projects", name, file))
	})
	return mux, nil
}

type Project struct {
	Name string
	Date string
}

// "GET /"
func projects_handler(w http.ResponseWriter, r *http.Request) {

	// Get project list
	entries, _ := os.ReadDir("projects")
	var projects_list []Project

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}

		info, _ := os.Stat(filepath.Join("projects", e.Name()))

		projects_list = append(projects_list, Project{
			Name: e.Name(),
			Date: info.ModTime().Format("02-Jan-2006"),
		})
	}

	// Render
	if utils.IsHTMX(r) {
		templ.Handler(projects(projects_list)).ServeHTTP(w, r)
		return
	}
	templ.Handler(ui.Layout(nil, ui.Nav_bar(), projects(projects_list))).ServeHTTP(w, r)
}

// "GET /{name}"
func show_project_handler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	templ.Handler(ui.Layout(nil, ui.Nav_bar(), project_canvas(name))).ServeHTTP(w, r)
}
