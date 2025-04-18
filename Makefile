VERSION := $(shell git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
REPO := $(shell basename -s .git `git config --get remote.origin.url`)

ifeq ($(BRANCH),main)
  FULL_VERSION := $(VERSION)
else ifeq ($(BRANCH),master)
  FULL_VERSION := $(VERSION)
else
  FULL_VERSION := $(VERSION)-$(REPO)
endif

.PHONY: help
help: ## print make targets 
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: download-tools
download-tools: ## download the required dependencies
	go run ./tools/downloader

.PHONY: build
build: ## build the application to the ./tmp folder
	go build -o ./tmp/gotth .

.PHONY: debug
debug: ## just debugging
	go run -ldflags "-X github.com/michielhemme/gotth/cmd.Version=$(VERSION)" . version