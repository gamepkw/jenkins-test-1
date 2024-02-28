package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (a *Handler) Get(c echo.Context) error {
	time.Sleep(1 * time.Second)

	fmt.Println("Get success")

	return c.JSON(http.StatusOK, Response{Message: "Get success"})
}
