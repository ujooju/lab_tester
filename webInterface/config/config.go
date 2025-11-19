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
	CurrentTaskOwner        string = ""
	CurrentTaskName         string = ""
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
		GiteaRedirectURI = os.Getenv("LT_GITEA_REDIRECT")
	}
	if GiteaRedirectURI == "" {
		return errors.New("missing LT_GITEA_REDIRECT setting")
	}

	if GiteaOauthCallbackState == "" {
		GiteaOauthCallbackState = os.Getenv("LT_GITEA_OAUTH_STATE")
	}
	if GiteaOauthCallbackState == "" {
		return errors.New("missing LT_GITEA_OAUTH_STATE setting")
	}

	if CurrentTaskOwner == "" {
		CurrentTaskOwner = os.Getenv("LT_CUR_TASK_OWNER")
	}
	if CurrentTaskOwner == "" {
		return errors.New("missing LT_CUR_TASK_OWNER setting")
	}

	if CurrentTaskName == "" {
		CurrentTaskName = os.Getenv("LT_CUR_TASK_NAME")
	}
	if CurrentTaskName == "" {
		return errors.New("missing LT_CUR_TASK_NAME setting")
	}

	return nil
}
