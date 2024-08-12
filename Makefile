ARCH = amd64
OS = linux

DOCKERFILE_PATH = ./deploy/docker/sso.Dockerfile
IMAGE_TARGET = release

IMAGE_NAME = go-sso-service
IMAGE_TAG  = latest
CONTAINER_NAME = sso-service

.PHONY: build run run-race
build:
	@CGO_ENABLED=0 GOARCH=${ARCH} GOOS=${OS} go build -o ./bin/sso ./cmd/sso

run:
	@CGO_ENABLED=0 GOARCH=${ARCH} GOOS=${OS} go run ./cmd/sso

run-race:
	@CGO_ENABLED=1 GOARCH=${ARCH} GOOS=${OS} go run -race ./cmd/sso

.PHONY: gen-proto gen-sqlc gen-all
gen-proto:
	$(foreach proto_file, $(shell find api/proto -name '*.proto'),\
		protoc --go_out=pb/ --go-grpc_out=pb/ \
		--proto_path=api/proto  $(proto_file);)

gen-sqlc:
	sqlc generate

gen-all: gen-proto gen-sqlc

.PHONY: docker-image docker-container
docker-image:
	@if [ ! -z $(docker images -q $(IMAGE_NAME):$(IMAGE_TAG))]; then docker image rm $(docker images -q $(IMAGE_NAME):$(IMAGE_TAG)); fi;
	docker build --tag $(IMAGE_NAME):$(IMAGE_TAG) --target $(IMAGE_TARGET) -f $(DOCKERFILE_PATH) . 

docker-container:
	docker run --rm -t -i \
		--name=$(CONTAINER_NAME) \
		--network=host \
		-d $(IMAGE_NAME):$(IMAGE_TAG)

run-test-db:
	docker compose up -d test_db migrator

stop-test-db:
	docker compose down test_db migrator

.PHONY: lint test 
lint:
	golangci-lint run --config .golangci.yaml ./...

test:
	@if [ -f coverage.txt ]; then rm coverage.txt; fi;
	@echo "Running unit tests with coverage profile"
	@go test ./... -coverprofile=coverage.txt -covermode=count
	@go tool cover -func=coverage.txt

.PHONY: clean
clean:
	rm ./bin/*
