.PHONY: help
help: ## print make targets 
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: download-tools
download-tools: ## download the required dependencies
	go run ./tools/downloader

.PHONY: build
build: ## build the application to the ./tmp folder
	go build -o ./tmp/gotth .