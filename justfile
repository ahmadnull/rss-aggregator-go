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
build: sqlc-gen sync
    go build -o ./target/rss-aggregator ./cmd/rss-aggregator

# Run rss-aggregator (default recipe)
run: goose-up build
    ./target/rss-aggregator

# Clean directory
clean:
    find -name '*~' -exec rm {} +

# Kill previous processes on port used by rss-aggregator
clean-port:
    kill $(lsof -wni tcp:$PORT | cut -d \n -f 2 | cut -d ' ' -f 2)

# Start pgadmin4 in detached tmux session and open in browser
pgadmin4:
    tmux new-session -d -s pgadmin4
    tmux send-keys "pgadmin4" C-m
    tmux detach -s pgadmin4 &2>/dev/null
    xdg-open http://127.0.0.1:5050/

# Kill pgadmin4 detached tmux session
kill-pgadmin4:
    tmux kill-session -t pgadmin4

# Goose postgres up migration
[working-directory: 'sql/schema']
goose-up:
    goose postgres $DB_URL up

# Goose postgres down migration
[working-directory: 'sql/schema']
goose-down:
    goose postgres $DB_URL down
 
# SQLC generate
sqlc-gen:
    sqlc generate

# Generate .env file
dotenv-gen PORT DB_URL:
    printf 'PORT=%s\nDB_URL=%s\n' {{PORT}} {{DB_URL}} | cat > .env
