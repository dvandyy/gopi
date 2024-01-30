include .env

# Start docker
run:
	docker-compose up $(args)

# Stop docker
stop:
	docker-compose down

# Clean rebuild
rebuild:
	docker-compose down --volumes && \
	docker-compose rm --all && \
	docker-compose pull --ignore-buildable && \
	docker-compose build --no-cache && \
	docker-compose up --force-recreate

# Build project to binary
build:
	go build -o ./bin/gopi ./cmd/

# Run migration up with golang-migrate
migrateup:
	migrate -path database/migration/ -database "postgresql://${DB_USER}:${DB_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" -verbose up

# Run migration down with golang-migrate
migratedown:
	migrate -path database/migration/ -database "postgresql://${DB_USER}:${DB_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" -verbose down

# Run tests
runtests:
	go test ./... -v

# Generate swagger files
swagger:
	docker exec gopi-dev sh & swag init --parseDependency -d ./api/v1/handlers -g ../../../cmd/main.go -o ./api/v1/docs

# Exec into main app
exec:
	docker exec -it gopi-dev sh

.PHONY: run stop rebuild build migrateup migratedown runtests swagger exec 