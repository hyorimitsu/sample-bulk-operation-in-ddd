package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/application/input"
	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/interface/presenter"
	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/registry"
)

type taskHandler struct {
	reg *registry.Registry
}

func (h taskHandler) ListTasks(ctx echo.Context) error {
	c := ctx.Request().Context()

	u := h.reg.TaskUseCaser
	dtos, err := u.ListTodos(c)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, presenter.MapToTasks(dtos))
}

func (h taskHandler) CreateTask(ctx echo.Context) error {
	c := ctx.Request().Context()

	var p input.TaskCreateParam
	if err := ctx.Bind(&p); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	u := h.reg.TaskUseCaser
	if err := u.CreateTodo(c, &p); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusNoContent, nil)
}

func (h taskHandler) UpdateTask(ctx echo.Context) error {
	c := ctx.Request().Context()

	id := ctx.Param("id")

	var p input.TaskUpdateParam
	if err := ctx.Bind(&p); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	u := h.reg.TaskUseCaser
	if err := u.UpdateTodo(c, id, &p); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusNoContent, nil)
}

func (h taskHandler) DeleteTask(ctx echo.Context) error {
	c := ctx.Request().Context()

	id := ctx.Param("id")

	u := h.reg.TaskUseCaser
	if err := u.DeleteTodo(c, id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
