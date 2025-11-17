package config

import (
	"errors"
	"os"
)

var (
	GiteaURL                string = ""
	GiteaClientID           string
	GiteaSecret             string
	GiteaRedirectURI        string = ""
	GiteaOauthCallbackState string = ""
)

func Confgure() error {
	if GiteaURL == "" {
		GiteaURL = os.Getenv("LT_GITEA_URL")
	}
	if GiteaURL == "" {
		return errors.New("missing LT_GITEA_URL setting")
	}

	if GiteaClientID == "" {
		GiteaClientID = os.Getenv("LT_GITEA_CLIENT_ID")
	}
	if GiteaClientID == "" {
		return errors.New("missing LT_GITEA_CLIENT_ID setting")
	}

	if GiteaSecret == "" {
		GiteaSecret = os.Getenv("LT_GITEA_SECRET")
	}
	if GiteaSecret == "" {
		return errors.New("missing LT_GITEA_SECRET setting")
	}

	if GiteaRedirectURI == "" {
		GiteaSecret = os.Getenv("LT_GITEA_REDIRECT")
	}
	if GiteaRedirectURI == "" {
		return errors.New("missing LT_GITEA_REDIRECT setting")
	}

	return nil
}
