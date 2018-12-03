# GALLY

## Configuation files

`.gally.yml`

```yml
strategies:
  compare-to:
    branch: master
  previous-commit:
    only: master
```

## Strategies

### compare-to

Test the updated projects between current branch and master

### previous-commit

On `master` branch test the updated projects the previous commit

- It works with git
- No nested project - at least for now
- Only one comparison strategy for now: compare files updated on current to master
