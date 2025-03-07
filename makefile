# Переменные для подключения к базе данных
DB_DSN := "postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Создание новой миграции
migrate-new:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir ./migrations "$${name}"

# Применение всех миграций
migrate-up:
	$(MIGRATE) up

# Применение одной миграции (если нужно пошаговое применение)
migrate-up-one:
	$(MIGRATE) up 1

# Откат последней миграции
migrate-down-one:
	$(MIGRATE) down 1

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
