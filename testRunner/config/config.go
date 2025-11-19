package config

import (
	"errors"
	"regexp"
)

var (
	GitURL          string = "https://infosec24.ru/git"
	GitURLProtoName string = "" //sets automatically
	GitURLHostName  string = "" //sets automatically
	ScriptName      string = "script.sh"
	ScriptLocation  string = "." //относительно main.go
	DatabaseURL     string = ""  //to be done
)

func Configure() error {
	reg := regexp.MustCompile("(.+)://(.+)")
	matches := reg.FindStringSubmatch("https://infosec24.ru/git")
	if len(matches) < 3 {
		return errors.New("invalid GitURL")
	}
	GitURLProtoName = matches[1]
	GitURLHostName = matches[2]
	return nil
}
