.PHONY: test help build watch tidy compose-runs compose-down compose-logs compose-restart
.DEFAULT_GOAL := help

test: ## Run the tests
	@go test -v ./...

tidy: ## tidy the go modules
	@echo "tidy"
	@go mod tidy
watch: ## Run the server with air
	@air -c .air.toml

build: ## Build the docker image
	@docker buildx build --platform linux/arm64 -t echo-app:latest .
compose-runs: ## Run docker compose
	@docker compose up -d
compose-down: ## Stop docker compose
	@docker compose down
compose-logs: ## Show docker compose logs
	@docker compose logs -f
compose-restart: ## Restart docker compose
	@docker compose restart
help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)