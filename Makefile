tag:
	@$(info tagging version v$(shell head -1 VERSION))
	@git tag v$(shell head -1 VERSION)
	@git push --tag