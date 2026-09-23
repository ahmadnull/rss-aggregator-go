# rss-aggregator-go

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
