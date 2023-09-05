package main

import (
	"log"
	"net/http"

	"github.com/MichalTaczala/go_course_building_modern_web_apps/pkg/config"
	"github.com/MichalTaczala/go_course_building_modern_web_apps/pkg/handlers"
	"github.com/MichalTaczala/go_course_building_modern_web_apps/pkg/render"
)

const portNumber = ":8070"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	app.TemplateCache = tc
	app.UseCache = false
	render.NewTemplates(&app)

	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
