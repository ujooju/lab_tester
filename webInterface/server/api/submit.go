package api

import (
	"net/http"

	"github.com/ujooju/lab_tester/webInterface/storage"
)

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	owner := r.FormValue("owner")
	name := r.FormValue("name")
	branch := r.FormValue("branch")
	if !HasAccess(r.Context().Value("token").(string), owner) {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}
	err := storage.SubmutTest(owner, name, branch)
	if err != nil {
		http.Error(w, "failed to submit", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/home/fork?owner="+owner+"&name="+name, http.StatusSeeOther)
}
