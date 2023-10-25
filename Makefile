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

## migrate: execute fresh migrations
.PHONY: migrate
migrate:
	@migrate -path db/migrations -database $(DB_DSN) up

## postgres: start postgres conitainer
.PHONY: postgres
postgres:
	@docker run -it -d -p 5432:5432 --name postgres -e POSTGRES_PASSWORD=postgres 2d74f8a2591ca34f2ea0565e385325570dd618e2d07372a8bb0ba97bddd19c93

## postgres/rm: clear container
.PHONY: postgres/rm
postgres/rm:
	@docker stop postgres
	@docker rm postgres

