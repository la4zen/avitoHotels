package main

import (
	"./routes"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/addRoom", routes.AddRoom)
	e.Logger.Fatal(e.Start(":8000"))
}
