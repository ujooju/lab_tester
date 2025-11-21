package handlers

import "net/http"

func ForkStatusHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/pages/fork.html")
}
