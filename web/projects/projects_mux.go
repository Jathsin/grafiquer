package projects

import (
	"jathsin/types"
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

	mux.HandleFunc("GET /{slug}", show_project_handler)

	mux.HandleFunc("GET /{slug}/static/{file...}", func(w http.ResponseWriter, r *http.Request) {
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

var metadata = types.Metadata{
	Title:                     "Projects",
	Meta_description:          "Interactive graphics experiments built with WebGL, shaders, and procedural systems. Explore visual simulations, noise generators, and generative graphics.",
	Meta_property_title:       "Graphics Projects — Grafiquer",
	Meta_property_description: "Explore interactive graphics experiments including procedural noise, shader filters, generative systems, and visual simulations.",
	Meta_Og_URL:               "https://grafiquer.com/projects",
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
	templ.Handler(ui.Layout(nil, ui.Nav_bar(), projects(projects_list), metadata)).ServeHTTP(w, r)
}

// "GET /{slug}"
func show_project_handler(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	// TODO: get specific metadata for project
	templ.Handler(ui.Layout(nil, ui.Nav_bar(), project_canvas(slug), metadata)).ServeHTTP(w, r)
}
