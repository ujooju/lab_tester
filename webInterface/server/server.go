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
	apiMux.HandleFunc("/api/list-forks", api.ListForksHandler)
	apiMux.HandleFunc("/api/fork-branches", api.ListBranchesHandler)
	apiMux.HandleFunc("/api/list-tests", api.ListTestsHandler)
	apiMux.HandleFunc("/api/submit", api.SubmitHandler)

	homeMux := http.NewServeMux()
	homeMux.HandleFunc("/home/", handlers.HomePageHandler)
	homeMux.HandleFunc("/home/fork", handlers.ForkStatusHandler)

	mux.Handle("/home/", middlewares.AuthMiddleware(homeMux))
	mux.Handle("/", loginMux)
	mux.Handle("/api/", middlewares.AuthMiddleware(apiMux))

	return mux
}
