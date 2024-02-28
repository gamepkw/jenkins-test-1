package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Transaction struct {
	Id          int64     `json:"id"`
	Amount      float64   `json:"amount"`
	Type        string    `json:"type"`
	Fee         float64   `json:"fee"`
	Total       float64   `json:"total"`
	SubmittedAt time.Time `json:"submitted_at"`
	CreatedAt   time.Time `json:"created_at"`
	Account     Account   `json:"account"`
	Receiver    Account   `json:"receiver,omitempty"`
}

type Account struct {
	AccountNo string     `json:"account_no,omitempty"`
	Uuid      string     `json:"uuid,omitempty"`
	Name      string     `json:"name,omitempty"`
	Email     string     `json:"email,omitempty"`
	Tel       string     `json:"tel,omitempty"`
	Balance   float64    `json:"balance"`
	Bank      string     `json:"bank,omitempty"`
	Status    string     `json:"status,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	IsClosed  int        `json:"is_closed,omitempty"`
}

func (a *TransactionHandler) GetTransaction(c echo.Context) error {
	time.Sleep(1 * time.Second)

	fmt.Println("Get all transaction history success")

	return c.JSON(http.StatusOK, TransactionResponse{Message: "Get all transaction history success"})
}
