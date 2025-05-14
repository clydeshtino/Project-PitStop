package routes

import (
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// TODO: Define function to show main page, call it here
	} else if r.Method == "POST" {
		// TODO: Define process function for main, call it here
	}
}
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

/*func routes() *mux.router {

}*/
