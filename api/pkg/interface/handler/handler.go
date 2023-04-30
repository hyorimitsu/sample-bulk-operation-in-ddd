package handler

import "github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/registry"

type Handler struct {
	taskHandler
}

func NewHandler(reg *registry.Registry) Handler {
	return Handler{
		taskHandler{reg: reg},
	}
}
