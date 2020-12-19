package routes

import (
	"github.com/labstack/echo"
)

func AddRoom(c echo.Context) error {
	return c.NoContent(200)
}
