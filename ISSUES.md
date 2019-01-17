# Known issues

The following file lists some known issues you might want to address to
improve `gally`. Some of those are questions and might come from a
misunderstanding of the tool.

## General

- When you type `gally` without any options, it displays
  `check subcommands` you might what to display the `inline` help.
- Why not put `version` and `build` as part of the `script` section?
- I do not like the `context` attribute as I find it confusing. I would prefer
  `workdir` for instance.
- Can you use more standard names for the project, like `darwin` (lowercase)
  and `x86_64` to ease the download from `uname` commands
- Could we have default values for the .gallu.yml file. For instance:
  - the name could match the directory name by default
  - the default strategies compare-to branch could be `master`
- WHat happens if you have no intersection with the `main` branch?

## Environment variables

- Could we name `GALLY_CWD`, `GALLY_VERSION`, `GALLY_NAME` with 
  `GALLY_PROJECT_WORKDIR`, `GALLY_PROJECT_VERSION` and `GALLY_PROJECT_NAME`
- Could we not add `GALLY_TAG`, rename it `GALLY_PROJECT_TAG` to the
  `gally list` command?
- Could we get `GALLY_PROJECT_ROOT` rather than `GALLY_CWD` and have also `GALLY_PROJECT_WORKDIR`
  anyway we need the path from the context and the path from the project itself.

# Diagnostics

- Could we check for the prerequisites in a command like `gally check`?
  - Are nested files detected?
  - Do you check if Git is installed?
- Could we not add a `dependencies` that would list the command to check
  project dependencies, e.g. `node --version`, `bash --version`...