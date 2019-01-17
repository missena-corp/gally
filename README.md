An opinionated tool to manage projects in a monorepository.

# Gally

<img align="right" src="https://user-images.githubusercontent.com/747/49454572-b0c3e600-f7e5-11e8-9be3-3feadfff1a52.jpeg" width="38%">

Gally helps to manage projects that are part of a monorepository. It provides
simple tools to detect what project have changed, as well as test and build
them.

> **Installation**: Download a binary from the [release page]

> **Requirement**: In order for Gally to work, you must have `git` installed
  and accessible from your path.

Gally is under the MIT license, see the [LICENSE](LICENSE) file for details.

## Configuration

To define a project with Gally, create a `.gally.yml` file in the directory
that conatins the project. For instance, if you have a project named
`simpleapi` in the `/apps/simpleapi` directory of your monorepository, create
a file `.gally.yml` in this directory.

> **Note**: Nested projects are not allowed

For details about the configuration, see:

- [Manifest](docs/MANIFEST.md) for details about the `.gally.yml` file
  properties
- [Environment variables](docs/VARIABLES.md) for the environment variables
  available from your scripts
- [Command Line Interface](docs/COMMAND.md) for details about how to run
  `gally`
- Using `gally` with [Continuous Integration](docs/CI.md) tools

## How are builds triggered?

Opposite to `scripts:` which are triggered if the project contains changes
the `build:` is triggered if the 2 following conditions are met:

- Changes are detected in the project
- A tag exists on the commit that matches the <project>@<version> and matches
  the `version:` command output.

> **Note**: We encourage you to rely on semver. As a result, we would suggest
  you tag your version of the `simpleapi` with `simpleapi@1.0.0` when you want
  to build the version for `1.0.0` assuming you have defined 1.0.0 in your
  version metadata.

## More

For more to come, see [TODO](TODO.md)

