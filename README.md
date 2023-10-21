# Mercedes Benz Arena Service

[![GoDoc](https://godoc.org/github.com/jo-hoe/mb-arena-service?status.svg)](https://godoc.org/github.com/jo-hoe/mb-arena-service)
[![Test Status](https://github.com/jo-hoe/mb-arena-service/workflows/test/badge.svg)](https://github.com/jo-hoe/mb-arena-service/actions?workflow=test)
[![Coverage Status](https://coveralls.io/repos/github/jo-hoe/mb-arena-service/badge.svg?branch=main)](https://coveralls.io/github/jo-hoe/mb-arena-service?branch=main)
[![Lint Status](https://github.com/jo-hoe/mb-arena-service/workflows/lint/badge.svg)](https://github.com/jo-hoe/mb-arena-service/actions?workflow=lint)
[![CodeQL Status](https://github.com/jo-hoe/mb-arena-service/workflows/CodeQL/badge.svg)](https://github.com/jo-hoe/mb-arena-service/actions?workflow=CodeQL)
[![Go Report Card](https://goreportcard.com/badge/github.com/jo-hoe/mb-arena-service)](https://goreportcard.com/report/github.com/jo-hoe/mb-arena-service)

Spiders the events from the mercedes benz arena in berlin and provides them via API.

> **Caution:** the API is still unstable and WILL change.

## Linting

Project used `golangci-lint` for linting.

### Installation

<https://golangci-lint.run/usage/install/>

### Execution

Run the linting locally by executing

```cli
golangci-lint run ./...
```

in the working directory

## ToDo

- add docu for
  - helm chart
  - services
  - how to release
- code
  - add tests for main
  - fix json property naming (lower case)
