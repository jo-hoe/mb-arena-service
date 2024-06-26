# Uber Arena Service

[![GoDoc](https://godoc.org/github.com/jo-hoe/mb-arena-service?status.svg)](https://godoc.org/github.com/jo-hoe/mb-arena-service)
[![Test Status](https://github.com/jo-hoe/mb-arena-service/workflows/test/badge.svg)](https://github.com/jo-hoe/mb-arena-service/actions?workflow=test)
[![Coverage Status](https://coveralls.io/repos/github/jo-hoe/mb-arena-service/badge.svg?branch=main)](https://coveralls.io/github/jo-hoe/mb-arena-service?branch=main)
[![Lint Status](https://github.com/jo-hoe/mb-arena-service/workflows/lint/badge.svg)](https://github.com/jo-hoe/mb-arena-service/actions?workflow=lint)
[![CodeQL Status](https://github.com/jo-hoe/mb-arena-service/workflows/CodeQL/badge.svg)](https://github.com/jo-hoe/mb-arena-service/actions?workflow=CodeQL)
[![Go Report Card](https://goreportcard.com/badge/github.com/jo-hoe/mb-arena-service)](https://goreportcard.com/report/github.com/jo-hoe/mb-arena-service)

Service that scrapes the [events from the Uber Arena website](https://www.uber-arena.de/events-tickets) and provides the details via API.

## Pre-requisites

- [golang](https://go.dev/doc/install) >= 1.22

## Execution

You can run the server just with Golang.

```bash
go run .
```

Or instead, use docker.

```bash
docker-compose up
```

## Configuration

The project can be configured with environment variables. The following variables are available.

| Key | Default | Description |
|-----|---------|-------------|
|CACHE_UPDATE_CRON|0 2 * * *|Event data is cached to lower the impact on the website. This setting describes how often the cache will be updated. The default is set to once a day at 02:00 AM|
|API_PORT|80|Port on with the API will listen|

## API Documentation

Provides on GET endpoint on a predefined port and will return a set of events.

```json
[
  {
    "name": "ALBA BERLIN - Armani Mailand",
    "link": "https://www.uber-arena.de/en/events/detail/alba-berlin-armani-mailand/2023-10-26-2000",
    "pictureUrl": "https://www.uber-arena.de/assets/img/ALBA-Mailand-f048ec885a.png",
    "start": "2023-10-26T20:00:00+02:00"
  },
  {
    "name": "Eisbären Berlin - Schwenninger Wild Wings",
    "link": "https://www.uber-arena.de/en/events/detail/eisbaeren-berlin-schwenningen/2023-10-27-1930",
    "pictureUrl": "https://www.uber-arena.de/assets/img/EBB-SWW-20a2dcc8b7.png",
    "start": "2023-10-27T19:30:00+02:00"
  },
  ...
]
```

## Linting

The project used `golangci-lint` for linting.

### Installation

<https://golangci-lint.run/welcome/install/>

### Run Linting

Run the linting locally by executing.

```cli
golangci-lint run ./...
```

in the working directory

## Future Work

- One can include also the Uber Eats Hall, as the parsing is largely the same
