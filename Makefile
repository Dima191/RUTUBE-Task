migrate_down:
	migrate --path ./migrations --database DATABASE_URL down

migrate_up:
	migrate --path ./migrations --database DATABASE_URL up

run:
	go build ./cmd/app
	./app

.PHONY: migrate_down, run

.DEFAULT_GOAL=run
