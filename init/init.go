package db

import (
	"errors"
	"log"

	"test-connect-db/configs"
	"test-connect-db/server/repository"
	"test-connect-db/server/repository/pg"
)

func Init(cfg *configs.Config) (*repository.Repo, error) {
	store := pg.NewPostgresStore(cfg)

	repo := pg.NewRepo(store.DB())
	if repo == nil {
		log.Fatal(errors.New("failed to create repo"))
	}

	return repo, nil
}
