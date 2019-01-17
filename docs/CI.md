`gally` is designed to ease the creation of CI workflows. It can be used with
different CI/CD tools.

# Using Gally with Travis-CI

Below is an example of a `.travis.yml` file that embeds `gally` to test/build
your monorepository projects with docker:

```yaml
sudo: required
language: node_js
node_js:
  - "10"
services:
  - docker
before_install:
  - wget https://github.com/missena-corp/gally/releases/download/v0.0.17/gally_0.0.17_linux_64-bit.tar.gz -O - | tar -zxvf - -C bin gally
script:
  - bin/gally run test
after_success:
  - bin/gally build
```
