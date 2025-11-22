package middlewares

import (
	"context"
	"net/http"

	"github.com/ujooju/lab_tester/webInterface/config"
	"github.com/ujooju/lab_tester/webInterface/storage"
)

// auth middleware
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("lt_user_id")
		if err != nil || cookie.Value == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		//val, ok := storage.Cache[cookie.Value]
		token, ok := storage.TokenCache.Get(cookie.Value)
		if !ok || token == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "token", token)))
	})
}

func AgentApiMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		token := r.FormValue("agent_secret")
		if token != config.AgentSecret {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// func ApiAuthMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		cookie, err := r.Cookie("lt_user_id")
// 		if err != nil || cookie.Value == "" {
// 			http.Error(w, "unathorized", http.StatusUnauthorized)
// 			return
// 		}
// 		val, ok := storage.Cache[cookie.Value]
// 		if !ok || val.Token == "" {
// 			http.Error(w, "unathorized", http.StatusUnauthorized)
// 			return
// 		}
// 		reqURL := config.GiteaURL + "/api/v1/user?access_token=" + val.Token
// 		authUserBytes, err := httpcurl.HttpCurl(httpcurl.CurlOption{
// 			"-X":         httpcurl.CurlValue{"GET"},
// 			"--location": httpcurl.CurlValue{reqURL},
// 			"-H":         httpcurl.CurlValue{"Content-Type: application/json"},
// 			"--tls-max":  httpcurl.CurlValue{"1.2"},
// 		}, time.Second*10)
// 		if err != nil || len(authUserBytes) == 0 {
// 			http.Error(w, "unathorized", http.StatusUnauthorized)
// 			return
// 		}
// 		var user gitea.User
// 		err = json.Unmarshal(authUserBytes, &user)
// 		if err != nil {
// 			http.Error(w, "unathorized", http.StatusUnauthorized)
// 			log.Println()
// 			return
// 		}
// 	})
// }
