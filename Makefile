.PHONY: help build watch
.DEFAULT_GOAL := help

tidy: ## tidy the go modules
	@echo "tidy"
	@go mod tidy
watch: ## Run the server with air
	@echo "watch"
	@air -c .air.toml

build: ## Build the docker image
	@docker buildx build --platform linux/arm64 -t echo-app:latest .
help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)