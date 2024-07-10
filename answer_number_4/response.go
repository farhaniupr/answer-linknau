package main

import (
	"github.com/labstack/echo/v4"
)

func ResponseInterface(c echo.Context, statusServer int, res interface{}, msg string) error {
	c.JSON(statusServer, JsonResponse{
		Status:   statusServer,
		Messages: msg,
		Data:     res,
	})
	return nil
}
