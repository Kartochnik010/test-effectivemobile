include .env

.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^//'


## run: run the cmd/api application
.PHONY: run
run:
	@go run ./cmd/api
