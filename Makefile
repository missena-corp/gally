.PHONY: goreleaser
goreleaser:
	@go get github.com/goreleaser/goreleaser && go install github.com/goreleaser/goreleaser

.PHONY: release
release: goreleaser
	@goreleaser --rm-dist

.PHONY: tag
tag:
	@$(info tagging version v$(shell head -1 VERSION))
	@git tag v$(shell head -1 VERSION)
	@git push --tag

.PHONY: test
test:
	@go test -race ./...
