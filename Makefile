SHELL := $(shell command -v bash;)
GO ?= go
# Static binary: runs in scratch/distroless images and cross-compiles without a C toolchain.
CGO_ENABLED ?= 0
GOOS ?= $(shell $(GO) env GOOS)
GOARCH ?= $(shell $(GO) env GOARCH)
GOCMD = CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO)

export PATH := $(shell $(GO) env GOPATH)/bin:$(PATH)

##@ Building

.PHONY: build-wowhabits
build-wowhabits: ## Build the wowhabits binary into bin/
	$(GOCMD) build \
		$(BUILDFLAGS) \
		-ldflags '$(LDFLAGS)' \
		-tags "$(BUILDTAGS)" \
		-o bin/wowhabits ./cmd/wowhabits

.PHONY: build-all
build-all: build-wowhabits ## Build all binaries

##@ Linting

.PHONY: lint-golangci-lint
lint-golangci-lint: ## Run `golangci-lint run ./...`
	golangci-lint run

.PHONY: lint-all
lint-all: lint-golangci-lint ## Run all linters

##@ Installation

.PHONY: install-golangci-lint
install-golangci-lint:  ## Install golangci-lint into "$GOPATH/bin"
	curl -sSfL --retry 5 https://golangci-lint.run/install.sh | sh -s -- -b $$($(GO) env GOPATH)/bin v2.14.0

.PHONY: install-all
install-all: install-golangci-lint ## Install all development tools

##@ Helpers

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_.-]+:.*?##/ { printf "  \033[36m%-21s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
