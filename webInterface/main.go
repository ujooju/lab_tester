package main

import (
	"log"

	"github.com/ujooju/lab_tester/webInterface/config"
	"github.com/ujooju/lab_tester/webInterface/server"
	"github.com/ujooju/lab_tester/webInterface/storage"
)

func main() {
	log.Println("Configuring...")
	err := config.Confgure()
	if err != nil {
		panic(err)
	}
	log.Println("Configured seccessfully")

	storage.InitCache()

	err = storage.InitSQLite()
	if err != nil {
		panic(err)
	}
	defer storage.DB.Close()

	server.Start()
}
