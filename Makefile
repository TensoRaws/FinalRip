GO ?= go

.DEFAULT_GOAL := default

version := v0.2.1
VS_PYTORCH_VERSION := v0.1.1

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
	docker buildx build -f ./deploy/worker-encode.dockerfile -t lychee0/finalrip-worker-encode --build-arg BASE_CONTAINER_TAG=cuda-${VS_PYTORCH_VERSION} .
	docker tag lychee0/finalrip-worker-encode lychee0/finalrip-worker-encode:latest
	docker tag lychee0/finalrip-worker-encode lychee0/finalrip-worker-encode:dev
	docker tag lychee0/finalrip-worker-encode lychee0/finalrip-worker-encode:cuda-dev
	docker tag lychee0/finalrip-worker-encode lychee0/finalrip-worker-encode:cuda-${version}
	docker tag lychee0/finalrip-worker-encode lychee0/finalrip-worker-encode:cuda

.PHONY: pt-release-dev
pt-release-dev:
	docker login
	docker push lychee0/finalrip-worker-encode:dev
	docker push lychee0/finalrip-worker-encode:cuda-dev

.PHONY: pt-release
pt-release:
	docker login
	docker push lychee0/finalrip-worker-encode:latest
	docker push lychee0/finalrip-worker-encode:cuda
	docker push lychee0/finalrip-worker-encode:cuda-${version}

.PHONY: pt-rocm
pt-rocm:
	docker buildx build -f ./deploy/worker-encode.dockerfile -t lychee0/finalrip-worker-encode --build-arg BASE_CONTAINER_TAG=rocm-${VS_PYTORCH_VERSION} .
	docker tag lychee0/finalrip-worker-encode lychee0/finalrip-worker-encode:rocm-dev
	docker tag lychee0/finalrip-worker-encode lychee0/finalrip-worker-encode:rocm-${version}
	docker tag lychee0/finalrip-worker-encode lychee0/finalrip-worker-encode:rocm

.PHONY: pt-rocm-release-dev
pt-rocm-release-dev:
	docker login
	docker push lychee0/finalrip-worker-encode:rocm-dev

.PHONY: pt-rocm-release
pt-rocm-release:
	docker login
	docker push lychee0/finalrip-worker-encode:rocm
	docker push lychee0/finalrip-worker-encode:rocm-${version}

.PHONY: release-dev
release-dev: pt pt-release-dev

.PHONY: release
release: pt pt-release

.PHONY: release-rocm-dev
release-rocm-dev: pt-rocm pt-rocm-release-dev

.PHONY: release-rocm
release-rocm: pt-rocm pt-rocm-release

.PHONY: dev
dev:
	docker compose -f ./deploy/docker-compose/lite/docker-compose.yml down
	docker compose -f ./deploy/docker-compose/lite/docker-compose.yml up -d
