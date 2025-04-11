run:
	go run ./cmd/server

tidy:
	go mod tidy

build:
	docker build -t go-backend-reuse .

down:
	docker-compose down

up:
	docker-compose up --build

ps:
	docker-compose ps

dev:
	kill -9 $(lsof -t -i :8090) && air -c air.toml

up-dev:
	docker-compose down -v
	docker-compose up --build

up-db:
	docker-compose up -d

reset-port:
	kill -9 $(lsof -t -i :8090)

reset:
	docker-compose down -v && docker-compose up -d