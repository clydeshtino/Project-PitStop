package controller

import (
	view "Project-Pitstop/src-app/shared/views"
	"net/http"
)

func ShowNews(w http.ResponseWriter, r *http.Request) {
	view.Render(w, "news.html", nil)
}

func ProcessNews(w http.ResponseWriter, r *http.Request) {

}

func ShowCompare(w http.ResponseWriter, r *http.Request) {
	view.Render(w, "compare.html", nil)
}

func ProcessCompare(w http.ResponseWriter, r *http.Request) {

}
