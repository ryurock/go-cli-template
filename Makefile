yq = docker run --rm -v "${PWD}":/workdir mikefarah/yq

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

.PHONY: release
release: ## リリースを行う
	@rm -rf artifacts
	@go build -o artifacts/MacOS/arm64/cli cmd/cli/main.go
	@cd artifacts/MacOS/arm64 && zip -r macOS-arm64 .

	@tag=`${yq} '.config.version' config/cli.yaml` && \
	git tag v$$tag && \
	git push origin v$$tag && \
	gh release create v$$tag "artifacts/MacOS/arm64/macOS-arm64.zip" \
	  --title="v$$tag" \
	  --notes-file CHANGELOG-template.md \
		--prerelease

.PHONY: release-rolback-rencent
release-rolback-rencent: ## リリースを行う
	@tag=`${yq} '.config.version' config/cli.yaml` && \
	git tag -d v$$tag || true && \
	git push --delete origin v$$tag || true && \
	gh release delete v$$tag || true
