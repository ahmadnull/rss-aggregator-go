set dotenv-load := true

_default: run

# Resolve dependencies
go-tidy:
    go mod tidy

# Download dependecies
go-get:
    go get

# Sync development environment
sync: go-tidy

# Build go module
build: sync
    go build -o ./target/rss-aggregator ./cmd/rss-aggregator

# Run rss-aggregator (default recipe)
run: build
    ./target/rss-aggregator

# Clean directory
clean:
    find -name '*~' -exec rm {} +

# Kill previous processes on port used by rss-aggregator
clean-port:
    kill $(lsof -wni tcp:$PORT | cut -d \n -f 2 | cut -d ' ' -f 2)
