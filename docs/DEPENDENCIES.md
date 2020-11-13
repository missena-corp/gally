# Dependencies

Dependencies in gally are related code, which, if updated, involves the code which depends on to be tested too.

## Declaration

Dependencies are declared in the config file as following:

```yml
dependencies:
  - ../../relative/path/to/dependance
  - /absolute/path/to/dependance
```

## Libraries

Some dependencies are libraries. They must have a `.gally.yml` file at the project root. And be declared as follow:

```yml
is_library: true
```

Being a library ensure it to be built before the project which depends on lauch a task
