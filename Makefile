env:
	@$(eval SHELL:=/bin/bash)
	@cp .env.example .env
	@echo "SECRET_KEY=$$(openssl rand -hex 32)" >> .env

dev:
	go mod download

tidy:
	go mod tidy

build-bin:
	go build ./cmd/main.go

run-bin:
	./main

run:
	go run ./cmd/main.go

up-dev:
	docker-compose -f dev.docker-compose.yml up -d

down-dev:
	docker-compose -f dev.docker-compose.yml down

build-prod:
	docker-compose -f prod.docker-compose.yml build

up-prod:
	docker-compose -f prod.docker-compose.yml up -d

down-prod:
	docker-compose -f prod.docker-compose.yml down

up-redis:
	docker run -d --name redis -p 6379:6379 redis:alpine

down-redis:
	docker stop redis

install-goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest

init-migrations:
	goose -dir migrations postgres "postgresql://postgres:2281337@localhost:5432/todo_app_db?sslmode=disablee" create init sql

migrate:
	goose -dir migrations postgres "postgresql://postgres:2281337@postgres:5432/todo_app_db?sslmode=disable" up
	# host=postgres if run in docker else host=localhost

run-prod:
	make migrate
	make run-bin
