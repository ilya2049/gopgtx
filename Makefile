MIGRATIONS_DIR=migrations
POSTGRES_DSN=host=127.0.0.1 port=5432 user=test-user password=password dbname=test-db sslmode=disable

.PHONY: run-read-committed
run-read-committed: clear-db
	go run cmd/app/*.go read_committed

.PHONY: run-repeatable-read
run-repeatable-read: clear-db
	go run cmd/app/*.go repeatable_read

.PHONY: run-serializable
run-serializable: clear-db
	go run cmd/app/*.go serializable

.PHONY: run-env
run-env:
	docker-compose -f deployments/docker-compose.yml up -d

.PHONY: stop-env
stop-env:
	docker-compose -f deployments/docker-compose.yml down -v 

.PHONY: clear-db
clear-db:
	goose -dir=${MIGRATIONS_DIR} postgres "${POSTGRES_DSN}" down
	goose -dir=${MIGRATIONS_DIR} postgres "${POSTGRES_DSN}" up