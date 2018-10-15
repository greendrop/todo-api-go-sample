package main

import (
	"github.com/greendrop/todo-api-go-sample/handler"
	"github.com/greendrop/todo-api-go-sample/model"
	"github.com/greendrop/todo-api-go-sample/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Validator
	e.Validator = &CustomValidator{validator: validator.New()}

	// DB
	service.InitDB()
	db := service.GetDBConnection()
	db.AutoMigrate(&model.Task{})

	// Routes
	api := e.Group("/api")
	{
		api.GET("/tasks", handler.GetTasks)
		api.GET("/tasks/:id", handler.GetTask)
		api.POST("/tasks", handler.CreateTask)
		api.PUT("/tasks/:id", handler.UpdateTask)
		api.DELETE("/tasks/:id", handler.DeleteTask)
	}

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
