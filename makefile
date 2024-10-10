PATH_RUN=cmd/server/main.go
GOOSE_DRIVER=postgres
GOOSE_DBSTRING=user=postgres password=123456789@ dbname=shopdevgo
GOOSE_MIGRATION_DIR=sql/schema

dev:
	go run ${PATH_RUN}
run:
	docker-compose up -d && go run ${PATH_RUN}
docker_build:
	docker-compose up -d --build
	docker-compose ps
docker_up:
	docker-compose up -d
docker_down:
	docker-compose down
docker_kill:
	docker-compose kill

upGoose:
	set "GOOSE_DRIVER=${GOOSE_DRIVER}" && set "GOOSE_DBSTRING=${GOOSE_DBSTRING}" && goose -dir ${GOOSE_MIGRATION_DIR} up

downGoose:
	set "GOOSE_DRIVER=${GOOSE_DRIVER}" && set "GOOSE_DBSTRING=${GOOSE_DBSTRING}" && goose -dir ${GOOSE_MIGRATION_DIR} down

resetGoose:
	set "GOOSE_DRIVER=${GOOSE_DRIVER}" && set "GOOSE_DBSTRING=${GOOSE_DBSTRING}" && goose -dir ${GOOSE_MIGRATION_DIR} reset

sqlgen:
	sqlc generate

.PHONY: run docker_build docker_up docker_down docker_kill
.PHONY: upGoose downGoose resetGoose