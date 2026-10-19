SHELL := /bin/bash

VERSION ?= $(shell git describe --tags --always --dirty)

##@ Build

.PHONY: binary
binary:  ## Build WoWHabits binary.
	@echo "INFO: Building WoWHabits binary..."
	docker buildx build \
		--build-arg VERSION=$(VERSION) \
		-f Dockerfile --target=result --output=. .

##@ Test

.PHONY: test
test: ## Run go unit test.
	go run gotest.tools/gotestsum@v1.13.0 --format pkgname -- ./...

##@ Lint

.PHONY: lint
lint:  ## Run golangci-lint.
	golangci-lint run

##@ Format

.PHONY: format
format: ## Format Go code with goimports (via golangci-lint).
	golangci-lint fmt

##@ Helpers

.PHONY: deps
deps: ## Tidy and verify Go module dependencies.
	go mod tidy
	go mod verify

.PHONY: clean
clean: ## Clean up build artifacts.
	rm -rf \
		$(wildcard coverage.*) \
		bin

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_.-]+:.*?##/ { printf "  \033[36m%-21s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
