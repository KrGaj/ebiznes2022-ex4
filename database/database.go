package database

import (
	"awesomeProject/database/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)
import "gorm.io/driver/sqlite"

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("sklep.db"), &gorm.Config{})

	if err != nil {
		panic("Doesn't work")
	}

	err = db.AutoMigrate(&models.User{}, &models.Category{}, &models.Product{}, &models.Purchase{}, &models.BoughtProduct{})

	if err != nil {
		panic(err)
	}

	return db
}

func DBMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	}
}
