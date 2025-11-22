package config

import (
	"errors"
	"os"
	"regexp"
)

var (
	GitURL          string       //TR_GIT_URL
	GitURLProtoName string       //sets automatically
	GitURLHostName  string       //sets automatically
	ScriptName      string       // TR_SCRIPT_NAME
	ScriptLocation  string = "." //относительно main.go //sets manually by hardcode))
	//DatabaseURL     string = ""
	CheckerName  string //gitea account with admin rights //TR_CHECKER_NAME
	CheckerToken string //pat for checher's account //TR_CHECKER_TOKEN
	LTURL        string //webInterface url //TR_LTURL
	AgentSecret  string //TR_AGENT_SECRET
	// agent secret to access webInterface
	// (sets mannually by env. var. LT_AGENT_SECRET for webInterface)
	// must be defined to access agent api
)

func Configure() error {
	if GitURL == "" {
		GitURL = os.Getenv("TR_GIT_URL")
	}
	if GitURL == "" {
		return errors.New("missing TR_GIT_URL setting")
	}

	reg := regexp.MustCompile("(.+)://(.+)")
	matches := reg.FindStringSubmatch("https://infosec24.ru/git")
	if len(matches) < 3 {
		return errors.New("invalid GitURL")
	}
	GitURLProtoName = matches[1]
	GitURLHostName = matches[2]

	if ScriptName == "" {
		ScriptName = os.Getenv("TR_SCRIPT_NAME")
	}
	if ScriptName == "" {
		return errors.New("missing TR_SCRIPT_NAME setting")
	}

	if CheckerName == "" {
		CheckerName = os.Getenv("TR_CHECKER_NAME")
	}
	if CheckerName == "" {
		return errors.New("missing TR_CHECKER_NAME setting")
	}

	if CheckerToken == "" {
		CheckerToken = os.Getenv("TR_CHECKER_TOKEN")
	}
	if CheckerToken == "" {
		return errors.New("missing TR_CKECKER_TOKEN setting")
	}

	if LTURL == "" {
		LTURL = os.Getenv("TR_LTURL")
	}
	if LTURL == "" {
		return errors.New("missing TR_LTURL setting")
	}

	if AgentSecret == "" {
		AgentSecret = os.Getenv("TR_AGENT_SECRET")
	}
	if AgentSecret == "" {
		return errors.New("missing TR_AGENT_SECRET setting")
	}

	return nil
}
