DB_URI="postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
# CMD
postgresup:
	@echo "Starting Docker Postgres Image"
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres
	@echo "Docker Postgres Image Started"
postgresdown:
	@echo "Stoping Docker Postgres"
	docker stop postgres
	@echo "Docker Postgres Stoped"

createdb:
	@echo "Creating simple_bank database"
	docker exec -it postgres createdb --username=root --owner=root simple_bank
	@echo "simple_bank database Created"

dropdb:
	@echo "Droping simple_bank database"
	docker exec -it postgres dropdb simple_bank
	@echo "simple_bank database Droped"

migrateup:
	@echo "Migrating UP database.."
	migrate -path db/migration -database ${DB_URI} -verbose up
	@echo "Migrated UP Successfuly"

migratedown:
	@echo "Migrating UP database.."
	migrate -path db/migration -database ${DB_URI} -verbose down
	@echo "Migrated UP Successfuly"

sqlc:
	@echo "Generating Sqlc code ..."
	sqlc generate
	@echo "Generated Successfuly"
test:
	go test -v -cover ./...
	
.PHONY: postgresup postgresdown createdb dropdb migrateup migratedown sqlc