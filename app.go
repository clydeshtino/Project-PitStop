package main

import (
	router "Project-Pitstop/src-app/routes"
	view "Project-Pitstop/src-app/shared/views"
	"embed"
	"log"
	"net/http"
)

var templates embed.FS
var resources embed.FS

func main() {
	log.Print("\nStarting Project-Pitstop.....")
	log.Print("\nProject-Pitstop is now running.")
	view.Parse(templates)
	router.SetResources(resources)

	// Using normal LandS for local testing, TLS will be incorporated later.
	log.Fatal(http.ListenAndServe(":8000", router.Routes()))
}
