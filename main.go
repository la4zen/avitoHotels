package main

import (
	_ "./routes"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/addRoom", routes.AddRoom)
	e.Logger.Fatal(e.Start(":8000"))
}
