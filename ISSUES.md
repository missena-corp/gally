# Known issues

The following file lists some known issues you might want to address to
improve gally. Some of those are questions and might come from a
misunderstanding of the tool:

- When you type `gally` without any options, it displays
  `check subcommands` you might what to display the `inline` help.
- Why not put `version` and `build` as part of the `script` section?
- Could we name GALLY_CWD, GALLY_VERSION, GALLY_NAME with 
  `GALLY_PROJECT_CWD`, `GALLY_PROJECT_VERSION` and `GALLY_PROJECT_NAME`
- Could we not add GALLY_TAG, rename it GALLY_PROJECT_TAG to the
  `gally list` command?
- Are nested files detected?
- Do you check if Git is installed?
- Could we get GALLY_PROJECT_ROOT rather than CWD and have also GALLY_PROJECT_CONTEXT
  anyway we need the path from the context and the path from the project itself.
