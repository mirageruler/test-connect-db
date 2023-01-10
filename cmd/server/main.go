package main

import (
	"fmt"
	"log"
	"net/http"

	"test-connect-db/configs"
	db "test-connect-db/init"
	"test-connect-db/server/api"
)

func main() {
	log.Println("Test connect db...")
	serve()
}

func serve() error {
	cfg, err := configs.New()
	if err != nil {
		log.Fatal(err)
	}

	repo, err := db.Init(cfg)
	if err != nil {
		log.Fatal(err)
	}

	ip := fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)
	srv := api.NewServer(repo)
	http.ListenAndServe(ip, srv)
	// block here

	return nil
}
