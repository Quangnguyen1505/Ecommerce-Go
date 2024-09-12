PATH_RUN=cmd/server/main.go

dev:
	go run ${PATH_RUN}
run:
	docker-compose up -d && go run ${PATH_RUN}
up:
	docker-compose up -d
down:
	docker-compose down
kill:
	docker-compose kill

.PHONY: run up down kill