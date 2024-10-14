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

.PHONY: lint ## pip install pre-commit
lint:
	golangci-lint run
	pre-commit install
	pre-commit run --all-files

.PHONY: all
all:
	docker buildx build -f ./deploy/server.dockerfile -t lychee0/finalrip-server .
	docker buildx build -f ./deploy/worker-cut.dockerfile -t lychee0/finalrip-worker-cut .
	docker buildx build -f ./deploy/worker-merge.dockerfile -t lychee0/finalrip-worker-merge .

.PHONY: pt
pt:
	docker buildx build -f ./deploy/worker-encode.dockerfile -t lychee0/finalrip-worker-encode --build-arg BASE_CONTAINER_TAG=cuda .
	docker tag lychee0/finalrip-worker-encode lychee0/finalrip-worker-encode:dev
	docker tag lychee0/finalrip-worker-encode lychee0/finalrip-worker-encode:cuda

.PHONY: pt-release
pt-release:
	docker login
	docker push lychee0/finalrip-worker-encode:dev
	docker push lychee0/finalrip-worker-encode:cuda

.PHONY: pt-rocm
pt-rocm:
	docker buildx build -f ./deploy/worker-encode.dockerfile -t lychee0/finalrip-worker-encode --build-arg BASE_CONTAINER_TAG=rocm .
	docker tag lychee0/finalrip-worker-encode lychee0/finalrip-worker-encode:rocm

.PHONY: pt-rocm-release
pt-rocm-release:
	docker login
	docker push lychee0/finalrip-worker-encode:rocm
