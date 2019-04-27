# TODO

The following file lists some known issues you might want to address to
improve `gally`:

- [ ] Find a way to avoid running multiple time tests for project sharing the same tests
- [ ] Add a `gally init` subcommand
- [x] Add `-f` option on run to bypass strategies
- [x] Add `-f` option on build not based on git tag
- [x] Add `-p` option to select project by name
- [ ] Add `--parallel n` to run commands in parallel
- [x] When you type `gally` without any options, it should display the help
- [ ] Change the `scripts:` content, see "Change in the script section" below
- [ ] Change the `context:` property with `workdir:`
- [ ] use more standard names for the project, like `darwin` (lowercase) and
      `x86_64` to ease the download from `uname` commands
- [ ] Add a `gally check` command to detect `.gally.yml` files in nested 
      directories AND check `git` is installed
- [ ] Add a `check:` properties with a list of commands, like `node --version`,
     `bash --version`...

- Could we have default values for the .gally.yml file. For instance:
  - the `name` could match the directory name by default
  - the default strategies compare-to branch could be `master`
- What happens if you have no intersection with the `main` branch?
- I would suggest we differentiate build which would be part of the `script:`
  section from the some kind of `push artefact` stage that would be part of the
  general section and is currently named `build`. I.e. I don't like the current
  naming

