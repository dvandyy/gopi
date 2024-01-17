include .env

docker-dev = docker-compose -f docker-compose-dev.yaml

# Start docker
run:
	$(docker-dev) up $(args)

# Stop docker
stop:
	$(docker-dev) down

# Clean rebuild
rebuild:
	$(docker-dev) down --volumes && \
	$(docker-dev) rm --all && \
	$(docker-dev) pull --ignore-buildable && \
	$(docker-dev) build --no-cache && \
	$(docker-dev) up --force-recreate

# Run migration up with golang-migrate
migrateup:
	migrate -path database/migration/ -database "postgresql://${DB_USER}:${DB_PASSWORD}@localhost:5432/gopi_db?sslmode=disable" -verbose up

# Run migration down with golang-migrate
migratedown:
	migrate -path database/migration/ -database "postgresql://${DB_USER}:${DB_PASSWORD}@localhost:5432/gopi_db?sslmode=disable" -verbose down

# Exec into main app
exec:
	docker exec -it gopi-dev sh

.PHONY: run stop rebuild migrateup migratedown exec