package server

import (
	"net/http"

	"github.com/ujooju/lab_tester/webInterface/server/handlers"
)

func Start() {
	mux := setMux()

	http.ListenAndServe(":3001", mux)
}

func setMux() *http.ServeMux {
	loginMux := http.NewServeMux()
	mux := http.NewServeMux()

	loginMux.HandleFunc("/login/gitea-oauth", handlers.GiteaOauthHandler)
	loginMux.HandleFunc("/", handlers.LoginHandler)

	mux.Handle("/", loginMux)
	return mux
}
