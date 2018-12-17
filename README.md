# Gally

<img align="right" src="https://user-images.githubusercontent.com/747/49454572-b0c3e600-f7e5-11e8-9be3-3feadfff1a52.jpeg" width="38%">

An opinionated tool to help with monoreposity projects.

It basicly helps us interfacing with `travis` avoiding to test and rebuild every projects on each change.
Each project contains a `.gally.yml` describing the project, and how to interact with it.

Constraints:

- It works with `git` command installed
- No nested project - at least for now
- 2 strategies checking files update (for now)

## Configuration files

They are named `.gally.yml`, and they look like the following example

```yml
name: example
ignore:
  - not/relevant/for/tests/*
scripts:
  build: docker build .
  test: echo hello world
strategies:
  compare-to:
    branch: master
  previous-commit:
    only: master
version: head -1 VERSION
build: echo go building!
```

They have to be placed in each managed projects.

## Strategies

Strategies are the way we check if files in project have been updated

### compare-to

Test the updated projects between current branch and master

### previous-commit

On `master` branch test the updated projects the previous commit

### Builds

In our workflow, final builds are launched with a specific tag. The tag is made
with the following schema: `project-name` + `@` + `semver version`. ie:

```
myproject@12.0.5
```

Builds are handled with the `build:` explaining how to run them.

### Special environment variables

- `GALLY_NAME`
- `GALLY_VERSION`

## TODO

- [x] Run scripts from project file
- [x] Automatically generate releases
- [x] Ability to ignore pattern
- [x] Handle project version
- [x] List projects
- [x] Add context dir option
- [x] Fix verbose flag
- [x] Handle `git`'s tag
- [x] Handle builds
- [ ] Remove `git`'s `.Exec` and do it through a library
- [ ] Find a way to avoid running multiple time tests for project sharing the same tests
- [ ] Add `gally init` subcommand
- [ ] Add `-f` option to bypass strategies
- [ ] Add `-p` option to select project by name
- [ ] Use channels to allow tests function `captureOutput` to run in parallel
- [ ] find a way to run `run` and `build` in parallel
