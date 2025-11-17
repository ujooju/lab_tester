package server

import (
	"net/http"

	"github.com/ujooju/lab_tester/testRunner/api"
)

func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/run", api.RunHandler)

	http.ListenAndServe(":80", mux)
}
