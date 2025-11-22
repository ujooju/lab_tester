package handlers

import (
	"log"
	"net/http"
	"net/url"

	"github.com/ujooju/lab_tester/webInterface/config"
)

func GiteaOauthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	values := url.Values{}
	values.Add("client_id", config.GiteaClientID)
	values.Add("redirect_uri", config.GiteaRedirectURI)
	values.Add("response_type", "code")
	values.Add("state", config.GiteaOauthCallbackState)
	http.Redirect(w, r, config.GiteaURL+"/login/oauth/authorize?"+values.Encode(), http.StatusSeeOther)
}
