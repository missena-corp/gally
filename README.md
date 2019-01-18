# Gally

<img align="right" src="https://user-images.githubusercontent.com/747/49454572-b0c3e600-f7e5-11e8-9be3-3feadfff1a52.jpeg" width="38%">

Gally helps to manage projects that are part of a monorepository. It provides
simple tools to detect what projects have changed, as well as test and build
them.

## Installation

Download a binary from the [release page](https://github.com/missena-corp/gally/releases)

**Requirement**: In order for Gally to work, you must have `git` installed
and accessible from your path.

## Configuration

To define a project with Gally, create a `.gally.yml` file in the directory
that contains the project. For instance, if you have a project named
`simpleapi` in the `/apps/simpleapi` directory of your monorepository, create
a file `.gally.yml` in this directory.

**Note**: Nested projects are not allowed

For more details about the configuration and how to use `gally`, see:

- [Manifest](docs/MANIFEST.md) for details about `.gally.yml` properties
- [Environment variables](docs/VARIABLES.md) for environment variables
  available from your scripts
- [Command Line Interface](docs/COMMAND.md) for details about how to run
  the `gally` command
- Using `gally` with [Continuous Integration](docs/CI.md) tools

## How are builds triggered?

Opposite to `scripts:` properties which are triggered if the project
contains changes the `build:` is triggered if the 2 following conditions
are met:

- Changes are detected in the project directories
- A tag exists on the commit that matches the <project>@<version> and matches
  the `version:` command output.

**Note**: We encourage you to rely on semver. As a result, we would suggest
you tag your version of the `simpleapi` with `simpleapi@1.0.0` when you want
to build the version for `1.0.0` assuming you have defined 1.0.0 in your
version metadata.

## More...

If you are interesting in what we are considering for what is coming next,
see [TODO](TODO.md). Do not hesitate to contribute by opening an issue or
create a pull request.

## License

Gally is available under the MIT license, see the [LICENSE](LICENSE) file for
details.

