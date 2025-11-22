package handlers

import (
	"log"
	"net/http"

	"github.com/ujooju/lab_tester/webInterface/storage"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("lt_user_id")
	if err != nil {
		http.Error(w, "failed to get cookies", http.StatusBadRequest)
		log.Println(err)
		return
	}
	if cookie.Value == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		log.Println(err)
		return
	}
	storage.TokenCache.Delete(cookie.Value)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
