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
	go test(srv)
	http.ListenAndServe(ip, srv)
	// block here

	return nil
}

func test(srv *api.Server) error {
	users, err := srv.Repo.Users.ManyByName("name-1")
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		return err
	}

	for i, u := range users {
		fmt.Printf("user#%d: %v", i+1, u)
	}

	return nil
}
