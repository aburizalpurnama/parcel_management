createdb:
	@echo creating database in docker container
	docker exec -it postgres14 createdb --username=postgres --owner=postgres parcel_management
	@echo success create database

dropdb:
	@echo deleting database in docker container
	docker exec -it postgres14 dropdb -U postgres parcel_management
	@echo success delete database

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
build_up:
	@echo Stopping docker images (if running...)
	docker-compose down
	@echo create all container and run it.
	docker compose up -d
	@echo success crete and run docker container

## down: stop docker compose
down:
	@echo Stopping docker compose...
	docker-compose down
	@echo Done!

## Migrate up
migrateup:
	@echo Migrate Up..
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/parcel_management?sslmode=disable" -verbose up
	@echo Success Migrate Up

## Migrate down
migratedown:
	@echo Migrate Down..
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/parcel_management?sslmode=disable" -verbose down
	@echo Success Migrate Down