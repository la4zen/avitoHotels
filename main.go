package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	"./routes"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		fmt.Errorf("db open error")
	}
	recipe, _ := ioutil.ReadFile("recipe.sql")
	db.Exec(string(recipe))
	db.Close()

	e := echo.New()

	e.POST("/addRoom", routes.AddRoom)
	e.POST("/delRoom", routes.DelRoom)
	e.POST("/getRooms", routes.GetRooms)

	e.POST("/addBooking", routes.AddBooking)
	e.POST("/delBooking", routes.DelBooking)
	e.POST("/getBooking", routes.GetBooking)

	e.Logger.Fatal(e.Start(":8000"))
}
