TAG := $(shell git describe --tags --exact-match 2>/dev/null)
VERSION := $(shell git tag --sort=-creatordate | head -n1 || echo "v0.0.0")
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

ifeq ($(TAG),)
  FULL_VERSION := $(VERSION)-$(BRANCH)
else
  FULL_VERSION := $(TAG)
endif

.PHONY: help
help: ## print make targets 
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: download-tools
download-tools: ## download the required dependencies
	go run ./tools/downloader

.PHONY: build
build: ## build the application to the ./tmp folder
	go build -ldflags "-X github.com/michielhemme/gotth/cmd.Version=$(FULL_VERSION)" -o ./tmp/gotth .

.PHONY: show-version
show-version: ## show current version of the application
	echo "Current version: $(FULL_VERSION)"