.PHONY: release
release:
	@curl -sL https://git.io/goreleaser | bash

.PHONY: tag
tag:
	@$(info tagging version v$(shell head -1 VERSION))
	@git tag v$(shell head -1 VERSION)
	@git push --tag

.PHONY: test
test:
	@go test -race ./...
