# go-ci-env

[![PkgGoDev](https://pkg.go.dev/badge/github.com/suzuki-shunsuke/go-ci-env/cienv)](https://pkg.go.dev/github.com/suzuki-shunsuke/go-ci-env/cienv)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/go-ci-env.svg)](https://github.com/suzuki-shunsuke/go-ci-env)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/go-ci-env/main/LICENSE)

Go library to get CI meta data from the environment variables.
go-ci-env supports various CI platform and provides unified API to abstract the difference of the platform.

## Supported CI services

* [AWS CodeBuild](https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-env-vars.html)
* [CircleCI](https://circleci.com/docs/2.0/env-vars/#built-in-environment-variables)
* [Drone](https://docs.drone.io/pipeline/environment/reference/)
* [GitHub Actions](https://docs.github.com/en/actions/configuring-and-managing-workflows/using-environment-variables#default-environment-variables)

## LICENSE

[MIT](LICENSE)
