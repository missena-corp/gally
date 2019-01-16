An opinionated tool to manage your projects in a monorepository.

# Gally

<img align="right" src="https://user-images.githubusercontent.com/747/49454572-b0c3e600-f7e5-11e8-9be3-3feadfff1a52.jpeg" width="38%">

Gally helps to manage projects that are part of a monorepository. It provides
simple tools to detect what project have changed, as well as test and build
them.

> Requirement: In order for Gally to work, you must have `git` installed and
  accessible from your path.

## Configuration

To define a project with Gally, you should create a `.gally.yml` file in the
project directory. For instance, if you have a project named `simpleapi` in
the `/apps/simpleapi` directory of your monorepository, create a file
`.gally.yml` in this directory.

> Note: Nested projects are not allowed

## Manifest parameters

The `.gally.yml` file contains a number of parameters that can be used to
customize the steps associated with the project build/test. Top level
properties are:

- `name` is the project name.
- `ignore` contains a list of subdirectories and files that can be
  modified without the project being considered as modified.

## Environment Variables

## Command Line

## Triggering Events

Tag <project>@<version> and the build is triggered on that.

## Example of Travis Contiguration

every projects on each change. Each project contains a `.gally.yml` file
describing the project, and how to interact with it.


- 2 strategies checking files update (for now)

## 

- Detects what project have changed in a branch

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

## Installing Gally with your CI

### Travis-CI


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

- `GALLY_ROOT` root path for the repository
- `GALLY_CWD` the 
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
- [ ] Add `-f` option on run to bypass strategies
- [ ] Add `-f` option on build not based on git tag
- [ ] Add `-p` option to select project by name
