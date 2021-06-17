package main

import (
	"UserCRUD/api"
	"github.com/labstack/echo"
)

/*
в файле db не совсем json,
там json объекты поочереди
но они не указаны как массив
мог проставить , и [] вначале и в конце, но решил так не делать
 */

func main() {
	println([]byte("\n"))
	e := echo.New()
	e.POST("/users", api.CreateUser)
	e.GET("/users/:id", api.GetUser)
	e.GET("/users", api.IndexUser)
	e.PUT("/users/:id", api.EditUser)
	e.DELETE("/users/:id", api.RemoveUser)

	e.Logger.Fatal(e.Start(":1323"))
}
