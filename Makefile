# Файл: Makefile
.PHONY: build run test clean docker-build docker-run docker-compose-up docker-compose-down migrate-up migrate-down

# Переменные
APP_NAME := digital-market
DOCKER_USERNAME := your-dockerhub-username

# Сборка приложения
build:
	go build -o $(APP_NAME) ./cmd/api

# Запуск приложения
run:
	go run ./cmd/api

# Запуск тестов
test:
	go test -v ./...

# Очистка бинарных файлов
clean:
	rm -f $(APP_NAME)

# Сборка Docker-образа
docker-build:
	docker build -t $(DOCKER_USERNAME)/$(APP_NAME):latest .

# Запуск Docker-контейнера
docker-run:
	docker run -p 8080:8080 $(DOCKER_USERNAME)/$(APP_NAME):latest

# Запуск с помощью Docker Compose
docker-compose-up:
	docker-compose up -d

# Остановка Docker Compose
docker-compose-down:
	docker-compose down

# Выполнение миграций вверх
migrate-up:
	migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/digital_market?sslmode=disable" up

# Откат миграций
migrate-down:
	migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/digital_market?sslmode=disable" down

# Создание новой миграции
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir ./migrations -seq $$name
