package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func newsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// TODO: Define function in controller to show news, call here
	} else if r.Method == "POST" {
		// TODO: Define function in controller to process news, then call here
	}
}

func compareHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// TODO: Define controller function showing compare then call
	} else if r.Method == "POST" {
		// TODO: Define process controller function, then call it here.
	}
}

// Router function to listen to our requests
func routes() *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/", newsHandler)
	router.HandleFunc("/compare", compareHandler)

	http.Handle("/", router)

	return router
}
