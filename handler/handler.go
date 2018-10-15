package handler

import (
	"github.com/greendrop/todo-api-go-sample/model"
	"github.com/greendrop/todo-api-go-sample/service"
	"github.com/labstack/echo"
	"net/http"
)

func GetTasks(c echo.Context) (err error) {
	db := service.GetDBConnection()
	tasks := []model.Task{}
	db.Find(&tasks)
	return c.JSON(http.StatusOK, tasks)
}

func GetTask(c echo.Context) (err error) {
	db := service.GetDBConnection()
	task := new(model.Task)
	db.First(&task, c.Param("id"))
	if db.NewRecord(task) {
		return c.JSON(http.StatusNotFound, map[string]string{})
	}
	return c.JSON(http.StatusOK, task)
}

func CreateTask(c echo.Context) (err error) {
	db := service.GetDBConnection()
	task := new(model.Task)
	if err := c.Bind(task); err != nil {
		return err
	}
	if err = c.Validate(task); err != nil {
		return
	}
	db.Create(&task)
	return c.JSON(http.StatusCreated, task)
}

func UpdateTask(c echo.Context) (err error) {
	db := service.GetDBConnection()
	task := new(model.Task)
	db.First(&task, c.Param("id"))
	if db.NewRecord(task) {
		return c.JSON(http.StatusNotFound, map[string]string{})
	}

	if err := c.Bind(task); err != nil {
		return err
	}
	if err = c.Validate(task); err != nil {
		return
	}
	db.Save(&task)
	return c.JSON(http.StatusNoContent, task)
}

func DeleteTask(c echo.Context) (err error) {
	db := service.GetDBConnection()
	task := new(model.Task)
	db.First(&task, c.Param("id"))
	if db.NewRecord(task) {
		return c.JSON(http.StatusNotFound, map[string]string{})
	}

	db.Delete(&task)
	return c.JSON(http.StatusNoContent, task)
}
