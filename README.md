# Gally <img align="right" src="https://user-images.githubusercontent.com/747/49454572-b0c3e600-f7e5-11e8-9be3-3feadfff1a52.jpeg" width=256>

An opinionated tool to help with monoreposity projects.
It basicly help us interfacing with `travis` avoiding to rebuild on each change.
Each project contains a `.gally.yml` describing the project, and how to interact with it.

Constraints:

- It works with `git` command installed
- No nested project - at least for now
- 2 strategies checking files update (for now)

## Configuation files

`.gally.yml`

```yml
name: example
scripts:
  build: docker build .
  test: echo hello world
strategies:
  compare-to:
    branch: master
  previous-commit:
    only: master
```

## Strategies

### compare-to

Test the updated projects between current branch and master

### previous-commit

On `master` branch test the updated projects the previous commit

## TODO

- [x] Run scripts from project file
- [ ] Add context dir option
- [ ] Remove `git`'s `.Exec` and do it through a library
- [ ] Handle `git`'s tag
- [ ] Handle project version
- [ ] Fix verbose flag
- [ ] Find a way to avoid running multiple time tests for project sharing the same tests
- [ ] Handle builds
- [ ] Automatically generate releases
- [ ] Ability to ignore pattern
