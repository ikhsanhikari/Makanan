package service

import (
	"Makanan/repository"
	"net/http"
	"github.com/labstack/echo"
)
func ReadAllUsers(c echo.Context) error {
		result := repository.ReadAll()
		return c.JSON(http.StatusOK, result)
}
