# `goose create <filename> sql` will create new migration file
.PHONY: up
up:
	goose -dir ./migrations postgres "${DATABASE}" up

.PHONY: down
down:
	goose -dir ./migrations postgres "${DATABASE}" down

.PHONY: status
status:
	goose -dir ./migrations postgres "${DATABASE}" status

.PHONY: run
run:
	go run cmd/main.go
