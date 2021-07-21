.PHONY:
.SILENT:
.DEFAULT_GOAL:= up

# ==============================================================================
# Main

init: full-clear docker-build docker-up \
 backend-init frontend-init \
 backend-ready frontend-ready
up: docker-up
down: docker-down
restart: down up
full-clear: docker-down-clear backend-clear frontend-clear
update-deps: frontend-yarn-upgrade backend-get-update restart

# ==============================================================================
# Docker support

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down --remove-orphans

docker-down-clear:
	docker-compose down -v --remove-orphans

docker-build:
	docker-compose build

docker-pull:
	docker-compose pull

# ==============================================================================
# Frontend commands

frontend-init: frontend-yarn-install

frontend-clear:
	docker run --rm -v ${PWD}/frontend:/app -w /app alpine sh -c 'rm -rf .ready dist'

frontend-yarn-install:
	docker-compose run --rm frontend-cli yarn install

frontend-yarn-upgrade:
	docker-compose run --rm frontend-cli yarn upgrade

frontend-ready:
	docker run --rm -v ${PWD}/frontend:/app -w /app alpine touch .ready

frontend-test-unit:
	docker-compose run --rm frontend-cli yarn test:unit

frontend-test-unit-watch:
	docker-compose run --rm frontend-cli yarn test:unit -- --watch

frontend-test-e2e:
	cd ./frontend/ && yarn test:e2e

frontend-cypress-install:
	cd ./frontend/ && ./node_modules/.bin/cypress install

# ==============================================================================
# Backend commands

backend-init: backend-install backend-wait-db migrate-up

backend-clear:
	docker run --rm -v ${PWD}/backend:/app -w /app alpine sh -c 'rm -rf .ready'

backend-install:
	docker-compose run --rm backend-cli go mod download
	docker-compose run --rm backend-cli go mod vendor

backend-get-update:
	docker-compose run --rm backend-cli go get -u -t -d -v ./...
	docker-compose run --rm backend-cli go mod tidy

backend-ready:
	docker run --rm -v ${PWD}/backend:/app -w /app alpine touch .ready

backend-wait-db:
	docker-compose run --rm backend-cli wait-for-it backend-postgres:5432 -t 30

backend-restart:
	docker-compose restart backend-golang
	docker-compose restart backend-cli

backend-test:
	docker-compose run --rm backend-cli go test --short -coverprofile=./var/test/cover.out -v ./...
	docker-compose run --rm backend-cli go tool cover -func=./var/test/cover.out

# ==============================================================================
# Go migrate postgresql

PSQL_URI=postgres://app:secret@backend-postgres:5432/app?sslmode=disable

migrate-create:
	docker-compose run --rm backend-cli migrate create -ext sql -dir migrations -seq $(f)

migrate-force:
	docker-compose run --rm backend-cli migrate -database $(PSQL_URI) -path migrations force 1

migrate-version:
	docker-compose run --rm backend-cli migrate -database $(PSQL_URI) -path migrations version

migrate-up:
	docker-compose run --rm backend-cli migrate -database $(PSQL_URI) -path migrations up 1

migrate-down:
	docker-compose run --rm backend-cli migrate -database $(PSQL_URI) -path migrations down 1

# ==============================================================================
# Tools commands

swaggo:
	echo 'Starting swagger generating'
	docker-compose run --rm backend-cli swag init -g cmd/api/main.go

deps-cleancache:
	docker-compose run --rm backend-cli go clean -modcache
