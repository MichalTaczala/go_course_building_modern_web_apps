package main

import (
	"log"
	"net/http"
	"time"

	"github.com/MichalTaczala/go_course_building_modern_web_apps/pkg/config"
	"github.com/MichalTaczala/go_course_building_modern_web_apps/pkg/handlers"
	"github.com/MichalTaczala/go_course_building_modern_web_apps/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8070"

var app config.AppConfig

var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction //true in production

	app.Session = session

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
