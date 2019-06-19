package delivery 


import (
	"Makanan/service"
	"github.com/labstack/echo"
)

func ReadAllUsersEndpoint() {
	e := echo.New()
	e.GET("/v1/api/user/readAll", service.ReadAllUsers)

	e.Logger.Fatal(e.Start(":1323"))
}