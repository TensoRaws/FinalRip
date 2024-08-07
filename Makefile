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

.PHONY: all
all:
	docker buildx build -f ./deploy/worker-cut.dockerfile -t lychee0/finalrip-worker-cut .
	docker buildx build -f ./deploy/worker-encode.dockerfile -t lychee0/finalrip-worker-encode .
	docker buildx build -f ./deploy/worker-merge.dockerfile -t lychee0/finalrip-worker-merge .
	docker buildx build -f ./deploy/server.dockerfile -t lychee0/finalrip-server .

.PHONY: pt
pt:
	docker buildx build -f ./deploy/worker-encode-pytorch.dockerfile -t lychee0/finalrip-worker-encode-pytorch .

.PHONY: pt-dev
pt-dev:
	docker buildx build -f ./deploy/worker-encode-pytorch.dockerfile -t lychee0/finalrip-worker-encode-pytorch .
	docker tag lychee0/finalrip-worker-encode-pytorch lychee0/finalrip-worker-encode-pytorch:dev
	docker login
	docker push lychee0/finalrip-worker-encode-pytorch:dev
