package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	ListTasks(ctx echo.Context) error
	CreateTask(ctx echo.Context) error
	UpdateTask(ctx echo.Context) error
	UpdateExpiredTaskToDone(ctx echo.Context) error
	BulkUpdateExpiredTasksToDone(ctx echo.Context) error
	DeleteTask(ctx echo.Context) error
}

func RegisterHandler(e *echo.Echo, hdl Handler, baseURL string) {
	e.GET(baseURL+"/healthz", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, nil)
	})

	e.GET(baseURL+"/tasks", hdl.ListTasks)
	e.POST(baseURL+"/tasks", hdl.CreateTask)
	e.PATCH(baseURL+"/tasks/:id", hdl.UpdateTask)
	e.PATCH(baseURL+"/expired_tasks/:id/done", hdl.UpdateExpiredTaskToDone)
	e.PATCH(baseURL+"/expired_tasks/done", hdl.BulkUpdateExpiredTasksToDone)
	e.DELETE(baseURL+"/tasks/:id", hdl.DeleteTask)
}
