package main

import (
	"github.com/ujooju/lab_tester/webInterface/config"
	"github.com/ujooju/lab_tester/webInterface/server"
)

func main() {
	err := config.Confgure()
	if err != nil {
		panic(err)
	}

	server.Start()
}
