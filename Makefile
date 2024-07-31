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

.PHONY: pt
pt:
	docker buildx build -f .\deploy\Dockerfile-worker-encode-pytorch -t lychee0/finalrip-worker-encode-pytorch .

.PHONY: pt-dev
pt-dev:
	docker buildx build -f .\deploy\Dockerfile-worker-encode-pytorch -t lychee0/finalrip-worker-encode-pytorch .
	docker tag lychee0/finalrip-worker-encode-pytorch lychee0/finalrip-worker-encode-pytorch:dev
	docker login
	docker push lychee0/finalrip-worker-encode-pytorch:dev

.PHONY: pt-release
pt-release:
	docker buildx build -f .\deploy\Dockerfile-worker-encode-pytorch -t lychee0/finalrip-worker-encode-pytorch .
	docker tag lychee0/finalrip-worker-encode-pytorch lychee0/finalrip-worker-encode-pytorch:release
	docker login
	docker push lychee0/finalrip-worker-encode-pytorch:release