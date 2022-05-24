package routes

import (
	"awesomeProject/database/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func ReadUser(c echo.Context) error {
	var id = c.QueryParam("ID")
	var db, _ = c.Get("db").(*gorm.DB)

	var user models.User
	db.Where("ID = ?", id).First(&user)

	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	var user models.User
	var bindErr = c.Bind(&user)

	if bindErr != nil {
		println(bindErr)
		return bindErr
	}

	println(user.ID)
	var db, _ = c.Get("db").(*gorm.DB)

	db.Create(&user)

	return c.JSON(http.StatusOK, user)
}
