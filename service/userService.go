package service

import (
	"Makanan/repository"
	"Makanan/model"
	"net/http"
	"github.com/labstack/echo"
	"strconv"
	"fmt"
)
func ReadAllUsers(c echo.Context) error {
		result := repository.ReadAll()
		return c.JSON(http.StatusOK, result)
}

func SaveUser(c echo.Context) error {
		Res := &model.ResponseModel{400,"Bad Request"}
		U := new(model.UserModel)
    	if err := c.Bind(U); err != nil {
        	return nil
    	}
		Res = repository.SaveUser(U)
		return c.JSON(http.StatusOK, Res)
}



func UpdateUser(c echo.Context) error {
		Res := &model.ResponseModel{400,"Bad Request"}
		U := new(model.UserModel)
    	if err := c.Bind(U); err != nil {
        	return nil
    	}
		Res = repository.UpdateUser(U)
		return c.JSON(http.StatusOK, Res)
}

func DeleteUser(c echo.Context) error {
		Res := &model.ResponseModel{400,"Bad Request"}
		id := c.QueryParam("id")
    	data,_ :=  strconv.Atoi(id)
    	fmt.Println("id", data)
		Res = repository.DeleteUser(data)
		return c.JSON(http.StatusOK, Res)
}