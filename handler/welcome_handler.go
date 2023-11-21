package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Welcome(c echo.Context) error {
	return c.HTML(http.StatusOK, `
		Welcome to Jougan which is written by Nim Team <br>
		To view metrics <a href="http://localhost:1994/metrics"><u>Click here: http://localhost:1994/metrics</u></a>
	`)
}
