package routes

import (
	"../models"
	"../util"
	"github.com/labstack/echo"
)

func AddRoom(c echo.Context) error {
	db := util.GetConnect()
	room := new(models.Room)
	c.Bind(room)
	return c.JSON(200, map[string]interface{}{})
}

func DelRoom(c echo.Context) error {
	room := new(models.Room)
	c.Bind(room)
	return c.NoContent(200)
}

func GetRooms(c echo.Context) error {
	var rooms []models.Room
	print(rooms)
	return c.NoContent(200)
}

func GetBooking(c echo.Context) error {
	booking := new(models.Booking)
	c.Bind(booking)
	return c.NoContent(200)
}
