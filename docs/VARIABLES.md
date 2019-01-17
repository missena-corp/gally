`gally` set environment variables that can be used in your scripts as well as
in the `.gally.yml` file:

- `GALLY_CWD` defines the project working directory, i.e. the location of the
  `.gally.yml` file with the to which you add the `context` property if it exists
- `GALLY_VERSION` defines the output from the `version:` command
- `GALLY_NAME` defines the project name
- `GALLY_ROOT` defines the root directory from gally, i.e. usually the repository
  top level directory
- `GALLY_TAG` if the commit is tagged with the project, defines the version of
  that tag.
