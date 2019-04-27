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

# Using Gally with Circle-CI

Below is an example of a `.circleci/config.yml` file that embeds `gally` to test/build
your monorepository projects with docker:

```yaml
version: 2.1
jobs:
  ci:
    working_directory: ~/project
    docker:
      - image: missena.io/circleci:latest
    steps:
      - add_ssh_keys:
          fingerprints:
            - "aa:aa:aa:aa:aa:aa:aa:aa:aa"
      - run:
          name: Clone repository
          command: |
            mkdir -p ~/.ssh
            ssh-keyscan -H github.com >> ~/.ssh/known_hosts
            git clone git@github.com:missena-corp/project.git ~/project
            git checkout -b ci $CIRCLE_SHA1
      - setup_remote_docker:
          docker_layer_caching: true
      - run:
          name: Tests
          command: |
            cd $HOME/project
            gally run test
      - run:
          name: Build
          command: |
            cd $HOME/project
            BUILD_BRANCH=$(grep -x "$CIRCLE_BRANCH" $HOME/ads/.ci-branches || true)
            if [[ -n "$BUILD_BRANCH" ]]; then
              echo "Building no tag projects"
              gally build
            fi

workflows:
  version: 2.1
  test_and_deploy:
    jobs:
      - ci:
          filters:
            branches:
              only: /.*/
```