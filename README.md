# rss-aggregator-go

[![CI workflow](https://img.shields.io/github/actions/workflow/status/ahmadnull/rss-aggregator-go/ci.yml?label=CI&logo=go)](https://github.com/ahmadnull/rss-aggregator-go/actions/workflows/ci.yml)

RSS Aggregator written in Go.

Thanks to: [Go Programming – Golang Course with Bonus Projects](https://www.youtube.com/watch?v=un6ZyFkqFKo) - [freeCodeCamp.org - YouTube](freeCodeCamp.org)

# Development

## Required (Build)

- go
- sqlc

## Required (Runtime)

- postgresql
- goose


## Recommended

- just
- pgadmin4

## Generate `.env` file

```sh
just dotenv-gen <PORT> <DB_URL> 
```
## Run

```sh
just run
```
