package server

import (
	"log"
	"net/http"

	"github.com/ujooju/lab_tester/webInterface/config"
	"github.com/ujooju/lab_tester/webInterface/server/api"
	"github.com/ujooju/lab_tester/webInterface/server/handlers"
	"github.com/ujooju/lab_tester/webInterface/server/middlewares"
)

func Start() {
	mux := setMux()

	log.Println("starting server at: ", config.Host+":"+config.Port)
	http.ListenAndServe(config.Host+":"+config.Port, mux)
}

func setMux() *http.ServeMux {
	loginMux := http.NewServeMux()
	mux := http.NewServeMux()

	loginMux.HandleFunc("/login/gitea-oauth", handlers.GiteaOauthHandler)
	loginMux.HandleFunc("/", handlers.LoginHandler)
	loginMux.HandleFunc("/logout", handlers.LogoutHandler)

	apiMux := http.NewServeMux()
	apiMux.HandleFunc("GET /api/list-forks", api.ListForksHandler)
	apiMux.HandleFunc("GET /api/fork-branches", api.ListBranchesHandler)
	apiMux.HandleFunc("GET /api/list-tests", api.ListTestsHandler)
	apiMux.HandleFunc("/api/submit", api.SubmitHandler)

	agentApiMux := http.NewServeMux()
	agentApiMux.HandleFunc("GET /agent/next-test", api.NextTestHandler)
	agentApiMux.HandleFunc("POST /agent/report", api.PostReportHandler)
	//apiMux.HandleFunc("GET /api/report", api.GetReportHandler)

	homeMux := http.NewServeMux()
	homeMux.HandleFunc("/home/", handlers.HomePageHandler)
	homeMux.HandleFunc("/home/fork", handlers.ForkStatusHandler)

	mux.Handle("/home/", middlewares.AuthMiddleware(homeMux))
	mux.Handle("/", loginMux)
	mux.Handle("/api/", middlewares.AuthMiddleware(apiMux))
	mux.Handle("/agent/", middlewares.AgentApiMiddleware(agentApiMux))

	return mux
}
