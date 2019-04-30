`gally` set environment variables that can be used in your scripts as well as
in the `.gally.yml` file:

- `GALLY_PROJECT_NAME` defines the project name
- `GALLY_PROJECT_WORKDIR` defines the project working directory, i.e. the
  location of the `.gally.yml` file with the to which you add the `context`
  property if it exists
- `GALLY_PROJECT_VERSION` defines the output from the `version:` command
- `GALLY_PROJECT_ROOT` defines the location of the project. This variable is
- `GALLY_ROOT` defines the root directory from gally, i.e. usually the repository
  top level directory
- `GALLY_TAG` if the commit is tagged with the project, defines the version of
  that tag.

**Important Note:** `GALLY_VERSION`, `GALLY_CWD` and `GALLY_NAME` are
deprecated, should not be used and will be remove in a soon to come
release.
