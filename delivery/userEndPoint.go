package main


import (
	"net/http"
	"Makanan/service"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/v1/api/user/readAll", func(c echo.Context) error {
		result := service.ReadAllUsers()
		return c.JSON(http.StatusOK, result)
	})
	e.Logger.Fatal(e.Start(":1323"))
}