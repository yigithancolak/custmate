GQL_GEN = github.com/99designs/gqlgen
DB_NAME = custmate
DB_USER = postgres
DB_PASS = secret
DB_HOST = localhost
DB_PORT = 5432
DB_URL = "postgresql://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"
MIGRATION_PATH = ./postgresdb/migration
VOLUME_PATH = ./volumes/postgres

graphql-generate:
	@go run $(GQL_GEN) generate

network-test:
	@sh -c 'docker network create $(DB_NAME)-test-net || true'

postgres-test: network-test
	@docker run --name $(DB_NAME)-test-db --network=$(DB_NAME)-test-net -p 5432:5432 -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASS) -e POSTGRES_DB=$(DB_NAME) -v $(VOLUME_PATH):/var/lib/postgresql/data -d postgres:12-alpine

pgadmin-test: network-test
	@docker run --rm --name $(DB_NAME)-test-pgadmin --network=$(DB_NAME)-test-net -p 5050:80 -e "PGADMIN_DEFAULT_EMAIL=admin@$(DB_NAME).com" -e "PGADMIN_DEFAULT_PASSWORD=secret" -d dpage/pgadmin4

migrateup:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose down

migratezero:
	migrate -path ${MIGRATION_PATH} -database "${DB_URL}" force 0

create-migration:
	@migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(name)

test:
	go test -v -cover -short ./...

.PHONY: graphql-generate network-test postgres-test pgadmin-test migrateup migratedown create-migration test
