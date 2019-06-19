package delivery 


import (
	"Makanan/service"
	"github.com/labstack/echo"
)

func UserEndpoint() {
	e := echo.New()
	e.GET("/v1/api/user/readAll", service.ReadAllUsers)
	e.POST("/v1/api/user/save", service.SaveUser)
	e.POST("/v1/api/user/update", service.UpdateUser)
	e.POST("/v1/api/user/delete", service.DeleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}