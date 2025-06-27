test:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out | grep total | awk '{print $$3}'

.PHONY: db-up migrate
db-up:
	docker compose up -d db

migrate:
	docker compose up migrate
