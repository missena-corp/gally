.PHONY: build
build:
	@go build -o dist/gally .

.PHONY: release
release:
	@curl -sL https://git.io/goreleaser | bash -s -- release --skip-validate --debug --rm-dist

.PHONY: tag
tag:
	@$(info tagging version v$(shell head -1 VERSION))
	@git tag v$(shell head -1 VERSION)
	@git push --tag

.PHONY: test
test:
	@go test -race ./...
