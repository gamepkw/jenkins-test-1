package handler

import (
	"github.com/labstack/echo/v4"
)

type Handler struct {
}

type Response struct {
	Message string `json:"message"`
}

func NewHandler(e *echo.Echo) {
	handler := &Handler{}

	e.GET("/get", handler.Get)
}
