.PHONY: run test build docker-build docker-run help

GO := go
DOCKER := docker

help:
	@echo "Available targets:"
	@echo "  make run           Run the app"
	@echo "  make test          Run Go tests"
	@echo "  make build         Build the binary"
	@echo "  make docker-build  Build the Docker image"
	@echo "  make docker-run    Run the Docker image"

run:
	$(GO) run ./cmd/receipt-manager

test:
	$(GO) test ./...

build:
	$(GO) build -o bin/receipt-manager ./cmd/receipt-manager

docker-build:
	$(DOCKER) build -t receipt-manager:local .

docker-run:
	$(DOCKER) run --rm -it \
		-e DISPLAY=$$DISPLAY \
		-v /tmp/.X11-unix:/tmp/.X11-unix:rw \
		--network host \
		receipt-manager:local
