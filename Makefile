include .env

## help: Print this message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^//'


## run: run the cmd/api application
.PHONY: run
run:
	@go run ./cmd/api

## migrate/up: migrations up
.PHONY: migrate/up
migrate/up:
	@migrate -path db/migrations -database $(DB_DSN) up

## migrate/down: migrations down
.PHONY: migrate/down
migrate/down:
	@migrate -path db/migrations -database $(DB_DSN) down

## postgres: start postgres conitainer
.PHONY: postgres
postgres:
	@docker run -it -d -p 5432:5432 --name postgres -e POSTGRES_PASSWORD=postgres 2d74f8a2591ca34f2ea0565e385325570dd618e2d07372a8bb0ba97bddd19c93

## postgres/rm: clear postgres container
.PHONY: postgres/rm
postgres/rm:
	@docker stop postgres
	@docker rm postgres

## postgres/psql: enter postgres container
.PHONY: postgres/psql
postgres/psql:
	@docker exec -it postgres bash
