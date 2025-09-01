GO ?= go

.DEFAULT_GOAL := default

.PHONY: tidy
tidy:
	${GO} mod tidy

.PHONY: server
server:
	${GO} build -o server ./server

.PHONY: worker
worker:
	${GO} build -o worker ./worker

.PHONY: test
test: tidy
	${GO} test ./...

.PHONY: lint
lint:
	golangci-lint run
	pre-commit install
	pre-commit run --all-files

.PHONY: all
all:
	docker buildx build -f ./deploy/server.dockerfile -t lychee0/finalrip-server .
	docker tag lychee0/finalrip-server lychee0/finalrip-server:dev

	docker buildx build -f ./deploy/worker-cut.dockerfile -t lychee0/finalrip-worker-cut .
	docker tag lychee0/finalrip-worker-cut lychee0/finalrip-worker-cut:dev

	docker buildx build -f ./deploy/worker-merge.dockerfile -t lychee0/finalrip-worker-merge .
	docker tag lychee0/finalrip-worker-merge lychee0/finalrip-worker-merge:dev

	docker buildx build -f ./deploy/dashboard.dockerfile -t lychee0/finalrip-dashboard .
	docker tag lychee0/finalrip-dashboard lychee0/finalrip-dashboard:dev

.PHONY: pt
pt:
	docker buildx build -f ./deploy/worker-encode.dockerfile -t lychee0/finalrip-worker-encode --build-arg BASE_CONTAINER_TAG=cuda .
	docker tag lychee0/finalrip-worker-encode lychee0/finalrip-worker-encode:dev
	docker tag lychee0/finalrip-worker-encode lychee0/finalrip-worker-encode:cuda-dev
	docker tag lychee0/finalrip-worker-encode lychee0/finalrip-worker-encode:cuda

.PHONY: dev
dev:
	docker compose -f ./deploy/docker-compose/lite/docker-compose.yml down
	docker compose -f ./deploy/docker-compose/lite/docker-compose.yml up -d
