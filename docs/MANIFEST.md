The `.gally.yml` file contains parameters that can be used to customize the
steps associated with the project build/test.

# Top level properties

Top level properties are:

- `build:` defines the command to build the project. This is not part
  of the `scripts` section because the triggering event 
- `workdir:` defines the working directory that should be use when running
  commands.
- `name:` is the project name; by default matches the directory name
- `ignore:` contains a list of subdirectories and files that can be
  modified without the project being considered as modified.
- `scripts:` contains the scripts to run `test` and other lifecyle
  commands
- `strategies:` defines how changes are checks, including, what branch
  should be considered as the main branch and when to compare with the
  previous commit compared to when compare with the base branch.
- `tag:` defines if builds rely on explicit `git tag` (true) or are
  automatically created when files change in the project (false). 
  Prefer the later!
- `version:` defines the command to figure out the project version

# Lifecycle (`scripts:`) properties

The `scripts`section of the file defines scripts associated with the
project lifecyle. For now, the only lifecycle step is the `test:` property
that defines the various commands that should run when running `gally run test`

# Change detection (`strategies:`) properties

`gally` behaves differently depending on the branch it relies. As a result,
it manages 2 different change detection strategies:

- It can compare changes with the previous commit. This is usually the strategy
  you apply to your main branch, e.g. `master`.
- It can compare changes with the last time it has been synchronize with your
  main branch.

Below is an example of a definition defining `master` as the main branch. With
this configuration, `gally` will compare the changes with the previous commit
if you are in the master branch. It will compare changes with the first

```yaml
strategies:
  compare-to:
    branch: master
  previous-commit:
    only: master
```

> **Limit and known issues**: The strategy above works well with a Github Flow
  as well as with Github `merge commits` and `squash merging`. However if you
  want to keep the commit from your branch when merging to your main branch,
  for instance with the github `rebase merging` or with a `git merge --ff`
  command, change will not be detected correctly. In addition if you are using
  a Git Flow, you might not be able to express the fact that you should compare
  to different branches.

# Example of a manifest file

Below is the example of a manifest file:

```yaml
name: simpleapi
workdir: .
ignore:
  - not/relevant/for/tests/*
scripts:
  test: make test
strategies:
  compare-to:
    branch: master
  previous-commit:
    only: master
version: git log -1 --format='%h' -- .
build: docker build .
tag: false
```
