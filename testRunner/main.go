package main

import (
	"github.com/ujooju/lab_tester/testRunner/config"
	"github.com/ujooju/lab_tester/testRunner/server"
)

func main() {
	err := config.Configure()
	if err != nil {
		panic(err)
	}
	server.StartServer()
}
