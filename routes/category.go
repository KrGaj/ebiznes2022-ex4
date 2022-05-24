package routes

import (
	"awesomeProject/database/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func ReadCategory(c echo.Context) error {
	var id = c.Param("id")
	var db, _ = c.Get("db").(*gorm.DB)
	var category models.Category
	db.Table("category").Where("ID = ?", id).First(&category)

	return c.JSON(http.StatusOK, category)
}
