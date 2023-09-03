package handlers

import (
	"net/http"

	"github.com/MichalTaczala/go_course_building_modern_web_apps/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")

}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")

}
