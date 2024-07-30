GO ?= go

.DEFAULT_GOAL := default

TAGS ?=

.PHONY: tidy
tidy: ## go mod tidy
	${GO} mod tidy

.PHONY: server
server: ## build binary file
	${GO} build -o server ./server

.PHONY: worker
worker: ## build binary file
	${GO} build -o worker ./worker

.PHONY: test
test: tidy ## go test
	${GO} test ./...

.PHONY: lint
lint:
	golangci-lint run
	pre-commit install # pip install pre-commit
	pre-commit run --all-files
