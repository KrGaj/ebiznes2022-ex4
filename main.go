package main

import (
	"awesomeProject/database"
	"awesomeProject/routes"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	var e = echo.New()
	var db = database.Connect()

	e.Use(database.DBMiddleware(db))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!")
	})

	e.GET("/users:id", routes.ReadUser)
	e.POST("/users", routes.CreateUser)

	e.Start(":2137")
}
