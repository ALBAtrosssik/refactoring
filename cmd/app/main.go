package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"petproject/internal/database"
	"petproject/internal/handlers"
	"petproject/internal/taskService"
	"petproject/internal/userService"
	"petproject/internal/web/tasks"
	"petproject/internal/web/users"
)

func main() {
	database.InitDB()

	taskRepo := taskService.NewTaskRepository(database.DB)
	taskService := taskService.NewTaskService(taskRepo)
	userRepo := userService.NewUserRepository(database.DB)
	userService := userService.NewUserService(userRepo)

	taskHandler := handlers.NewTaskHandler(taskService)
	userHandler := handlers.NewUserHandler(userService)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictTaskHandler := tasks.NewStrictHandler(taskHandler, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictTaskHandler)
	strictUserHandler := users.NewStrictHandler(userHandler, nil) // тут будет ошибка
	users.RegisterHandlers(e, strictUserHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
