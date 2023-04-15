.PHONY: help
help: ## make taskの説明を表示する
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: install
install: ## 初期構築を行う
	docker compose -f build/docker/docker-compose.yaml build

.PHONY: login
login: ## コンテナにログインする
	docker compose -f build/docker/docker-compose.yaml run --rm app /bin/ash

.PHONY: build
build: ## ビルドを行う
	docker compose -f build/docker/docker-compose.yaml run --rm app go build -o build/bin/ cmd/cli.go