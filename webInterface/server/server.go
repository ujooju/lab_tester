package server

import (
	"net/http"

	"github.com/ujooju/lab_tester/webInterface/server/api"
	"github.com/ujooju/lab_tester/webInterface/server/handlers"
	"github.com/ujooju/lab_tester/webInterface/server/middlewares"
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
	loginMux.HandleFunc("/logout", handlers.LogoutHandler)

	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/api/task-forks", api.TaskForksHandler)

	homeMux := http.NewServeMux()
	homeMux.HandleFunc("/home", handlers.HomePageHandler)

	mux.Handle("/home", middlewares.AuthMiddleware(homeMux))
	mux.Handle("/", loginMux)
	mux.Handle("/api/", middlewares.AuthMiddleware(apiMux))

	return mux
}
