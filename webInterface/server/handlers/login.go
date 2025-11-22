package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	httpcurl "github.com/ujooju/http-curl/lib"
	"github.com/ujooju/lab_tester/webInterface/config"
	"github.com/ujooju/lab_tester/webInterface/storage"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		GiteaOauthHandler(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cookie, err := r.Cookie("lt_user_id")
	if err == nil {
		log.Println(err)
		if _, ok := storage.TokenCache.Get(cookie.Value); ok {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}
	}

	fmt.Println(w, r.URL.RawQuery)
	state := r.URL.Query().Get("state")
	if state == "" {
		http.ServeFile(w, r, "static/pages/login.html")
		return
	}

	if state != config.GiteaOauthCallbackState {
		http.Error(w, "invalid state", http.StatusBadRequest)
		return
	}

	//fmt.Fprintf(w, "OAuth callback received with state: %s\n", state)

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "missing code", http.StatusBadRequest)
		return
	}

	accessRequest := GiteaAccessTokenRequest{
		ClientID:     config.GiteaClientID,
		ClientSecret: config.GiteaSecret,
		GrantType:    "authorization_code",
		Code:         code,
		RedirectURI:  config.GiteaRedirectURI,
	}

	accessRequestJSONBytes, err := json.Marshal(accessRequest)
	if err != nil {
		http.Error(w, "failed to marshal access request", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	response, err := httpcurl.HttpCurl(httpcurl.CurlOption{
		"-X":         httpcurl.CurlValue{"POST"},
		"-d":         httpcurl.CurlValue{string(accessRequestJSONBytes)},
		"--location": httpcurl.CurlValue{config.GiteaURL + "/login/oauth/access_token"},
		"-H":         httpcurl.CurlValue{"Content-Type: application/json"},
		"--tls-max":  httpcurl.CurlValue{"1.2"},
	}, time.Second*10)

	if err != nil {
		http.Error(w, "failed to get access token", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var accessTokenResponse GiteaAccessTokenResponse
	err = json.Unmarshal(response, &accessTokenResponse)
	if err != nil {
		http.Error(w, "failed to unmarshal access token response", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	userID := uuid.New().String()
	storage.TokenCache.Set(userID, accessTokenResponse.AccessToken)
	http.SetCookie(w, &http.Cookie{
		Name:  "lt_user_id",
		Value: userID,
	})
	//fmt.Fprintln(w, userID)
	//fmt.Fprintln(w, string(response))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

type GiteaAccessTokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
	RedirectURI  string `json:"redirect_uri"`
}

type GiteaAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}
