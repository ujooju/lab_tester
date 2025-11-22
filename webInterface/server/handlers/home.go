package handlers

import "net/http"

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/pages/home.html")
}
