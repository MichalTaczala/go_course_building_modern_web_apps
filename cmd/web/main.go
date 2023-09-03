package main

import (
	"net/http"

	"github.com/MichalTaczala/go_course_building_modern_web_apps/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	_ = http.ListenAndServe(portNumber, nil)

}
