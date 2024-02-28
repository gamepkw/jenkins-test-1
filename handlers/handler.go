package handler

import (
	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
}

type TransactionResponse struct {
	Message string       `json:"message"`
	Body    *Transaction `json:"body,omitempty"`
}

func NewTransactionHandler(e *echo.Echo) {
	handler := &TransactionHandler{}

	e.GET("/get-all-transaction", handler.GetTransaction)
}
