.PHONY: run
run:
	go run cmd/server/run.go

.PHONY: migrate-create
migrate-create:
	goose -dir=./server/repository/pg/migrations create migration sql

.PHONY: db-init
db-init:
	docker-compose down && docker-compose up -d test_database

