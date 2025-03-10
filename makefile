# Переменные для подключения к базе данных
DB_DSN := "postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Создание новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations "$(NAME)"

# Применение всех миграций
migrate-up:
	$(MIGRATE) up

# Полный откат всех миграций (использовать осторожно!)
migrate-down:
	$(MIGRATE) down

# Генерация API-кода для tasks
gen-tasks:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go

# Генерация API-кода для users
gen-users:
	oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/openapi.yaml > ./internal/web/users/api.gen.go

# Запуск линтера
lint:
	golangci-lint run --out-format=colored-line-number

# Запуск приложения
run:
	go run cmd/app/main.go
