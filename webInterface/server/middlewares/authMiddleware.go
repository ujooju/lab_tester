package middlewares

import (
	"context"
	"net/http"

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
		val, ok := storage.Cache[cookie.Value]
		if !ok || val.Token == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "token", val.Token)))
	})
}
