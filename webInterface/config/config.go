package config

import (
	"errors"
	"log"
	"os"
	"strings"
)

var (
	Host                    string //optional. by default: 127.0.0.1
	Port                    string
	GiteaURL                string
	GiteaClientID           string //oauth
	GiteaSecret             string //oauth
	GiteaRedirectURI        string //oauth
	GiteaOauthCallbackState string //oauth
	CurrentTaskOwner        string
	CurrentTaskName         string
	Admins                  []string = []string{} //f.ex.: LT_ADMINS=gitt,ujooju,admin
	AgentSecret             string                //must be set manually with env. variable
)

func Confgure() error {
	if Port == "" {
		Port = os.Getenv("LT_PORT")
	}
	if Port == "" {
		return errors.New("missing LT_PORT setting")
	}

	if Host == "" {
		Host = os.Getenv("LT_HOST")
	}
	if Host == "" {
		Host = "127.0.0.1"
		log.Println("using host 127.0.0.1")
	}

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

	if len(Admins) == 0 {
		adminsStr := os.Getenv("LT_ADMINS")
		Admins = strings.Split(adminsStr, ",")
	}
	if len(Admins) == 0 {
		return errors.New("missing LT_ADMINS setting")
	}

	if AgentSecret == "" {
		AgentSecret = os.Getenv("LT_AGENT_SECRET")
	}
	if AgentSecret == "" {
		return errors.New("missing LT_AGENT_SECRET setting")
	}

	return nil
}
