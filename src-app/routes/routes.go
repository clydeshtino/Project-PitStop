package routes

import (
	controller "Project-Pitstop/src-app/controllers"
	"embed"
	"net/http"

	"github.com/gorilla/mux"
)

var resources embed.FS

func SetResources(fs embed.FS) {
	resources = fs
}

func NewsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		controller.ShowNews(w, r)
	} else if r.Method == "POST" {
		controller.ProcessNews(w, r)
	}
}

func CompareHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		controller.ShowCompare(w, r)
	} else if r.Method == "POST" {
		controller.ProcessCompare(w, r)
	}
}

// Router function to listen to our requests
func Routes() *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/", NewsHandler)
	router.HandleFunc("/compare", CompareHandler)

	http.Handle("/", router)

	return router
}
