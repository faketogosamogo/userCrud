package api

import (
	"UserCRUD/infrastructure"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"math"
	"net/http"
	"strconv"
)
var Context = infrastructure.ContextNew()

func GetUser(c echo.Context) error {
	id := c.Param("id")
	user:= Context.GetUser(id)
	if user == nil{
		return c.String(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	u := new(UserCreateVm)
	err := c.Bind(u)
	if err!=nil{
		return c.String(http.StatusUnprocessableEntity, "invalid request body")
	}
	newUser := infrastructure.User{
		Id:   uuid.New().String(),
		Name: u.Name,
	}
	Context.AddUser(&newUser)
	return c.JSON(http.StatusCreated, newUser)
}

func EditUser(c echo.Context) error {
	id := c.Param("id")
	u := new(UserCreateVm)
	err := c.Bind(u)
	if err!=nil{
		return c.String(http.StatusUnprocessableEntity, "invalid request body")
	}
	user:= infrastructure.User{
		Id: id,
	 	Name:  u.Name,
	}
	if Context.EditUser(&user) == nil{
		return c.String(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, user)
}

func RemoveUser(c echo.Context) error {
	id := c.Param("id")
	user:= Context.GetUser(id)
	if user == nil{
		return c.String(http.StatusNotFound, "user not found")
	}
	Context.RemoveUser(user)
	return c.JSON(http.StatusOK, user)
}

func IndexUser(c echo.Context) error{
	page, _ := strconv.Atoi(c.QueryParam("page"))
	perPage, _ := strconv.Atoi(c.QueryParam("perPage"))
	if page < 1{
		page = 1
	}
	if perPage < 1{
		perPage = 8
	}
	users:=Context.IndexUser(perPage, page)
	countOfRecords := Context.CountOfRecords()

	c.Response().Header().Add("x-total-pages", strconv.Itoa(totalPages(perPage, countOfRecords)))
	return c.JSON(http.StatusOK, users)
}

func totalPages(perPage, countOfRecords int) int {
	if countOfRecords==0{
		return 0
	}
	temp := float64(countOfRecords) / float64(perPage)
	return int(math.Ceil(temp))
}