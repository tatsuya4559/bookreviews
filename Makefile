# `goose create <filename> sql` will create new migration file
.PHONY: up
up:
	goose -dir ./migrations sqlite3 ./dev.db up

.PHONY: down
down:
	goose -dir ./migrations sqlite3 ./dev.db down

.PHONY: status
status:
	goose -dir ./migrations sqlite3 ./dev.db status

.PHONY: run
run:
	go run cmd/main.go
