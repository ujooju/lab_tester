package api

import (
	"encoding/json"
	"net/http"
	"slices"
	"time"

	"code.gitea.io/sdk/gitea"
	httpcurl "github.com/ujooju/http-curl/lib"
	"github.com/ujooju/lab_tester/webInterface/config"
	"github.com/ujooju/lab_tester/webInterface/storage"
)

func HasAccess(token string, userName string) bool {
	reqURL := config.GiteaURL + "/api/v1/user?access_token=" + token
	authUserBytes, err := httpcurl.HttpCurl(httpcurl.CurlOption{
		"-X":         httpcurl.CurlValue{"GET"},
		"--location": httpcurl.CurlValue{reqURL},
		"-H":         httpcurl.CurlValue{"Content-Type: application/json"},
		"--tls-max":  httpcurl.CurlValue{"1.2"},
	}, time.Second*10)
	if err != nil || len(authUserBytes) == 0 {
		return false
	}
	var user gitea.User
	err = json.Unmarshal(authUserBytes, &user)
	if err != nil {
		return false
	}
	if user.UserName == userName || slices.Contains(config.Admins, user.UserName) {
		return true
	}
	return false
}

func ListTestsHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	owner := r.FormValue("owner")
	name := r.FormValue("name")
	if !HasAccess(r.Context().Value("token").(string), owner) {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}
	testRecords, err := storage.GetTestsByOwnerAndName(owner, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	testRecordsBytes, err := json.Marshal(testRecords)
	if err != nil {
		http.Error(w, "failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Write(testRecordsBytes)
}
