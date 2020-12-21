package routes

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/araddon/dateparse"

	"../models"
	"github.com/labstack/echo"
)

func AddRoom(c echo.Context) error {
	var room models.Room
	c.Bind(&room)
	if room.Price == nil {
		return c.String(http.StatusBadRequest, "price requeired")
	}
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return c.String(http.StatusInternalServerError, "db connection failed")
	}
	defer db.Close()
	row, err := db.Exec("INSERT INTO rooms(discription, price) VALUES(?,?)", room.Discription, room.Price)
	if err != nil {
		return c.String(http.StatusInternalServerError, "insert room error")
	}
	id, _ := row.LastInsertId()
	return c.JSON(200, map[string]interface{}{
		"room_id": id,
	})
}

func DelRoom(c echo.Context) error {
	room := new(models.Room)
	c.Bind(&room)
	if room.ID == nil {
		return c.String(400, "Room id required")
	}
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return c.String(http.StatusInternalServerError, "db connection failed")
	}
	defer db.Close()
	db.Exec("DELETE FROM rooms WHERE id = ?; DELETE FROM booking WHERE roomid = ?", room.ID, room.ID)
	return c.NoContent(200)
}

func GetRooms(c echo.Context) error {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return c.String(500, "db connection failed")
	}
	defer db.Close()
	var rooms []models.Room
	rows, err := db.Query("SELECT * FROM rooms")
	if err != nil {
		return c.String(500, "get rows error")
	}
	defer rows.Close()
	for rows.Next() {
		var room models.Room
		err := rows.Scan(&room.ID, &room.Discription, &room.Price, &room.CreationDate)
		if err != nil {
			return c.String(500, "scan row failed")
		}
		rooms = append(rooms, room)
	}
	return c.JSON(200, rooms)
}

func AddBooking(c echo.Context) error {
	var book struct {
		ID        *int    `json:"id"`
		DateStart *string `json:"date_start"`
		DateEnd   *string `json:"date_end"`
	}
	c.Bind(&book)
	if book.ID == nil || book.DateStart == nil || book.DateEnd == nil {
		c.String(400, "bookid, datestart, dateend required")
	}
	t1, err := dateparse.ParseAny(*book.DateStart)
	if err != nil {
		return c.String(400, "invalid datestart")
	}
	t2, err := dateparse.ParseAny(*book.DateEnd)
	if err != nil {
		return c.String(400, "invalid dateend")
	}
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return c.String(500, "db connection failed")
	}
	defer db.Close()
	result, err := db.Exec("INSERT INTO booking(roomid, datestart, dateend) VALUES (?,?,?)", book.ID, t1, t2)
	if err != nil {
		return c.String(500, "db insert error")
	}
	id, _ := result.LastInsertId()
	return c.JSON(200, map[string]interface{}{
		"booking_id": id,
	})
}

func DelBooking(c echo.Context) error {
	var booking models.Booking
	c.Bind(&booking)
	if booking.ID == nil {
		return c.String(400, "booking id required")
	}
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return c.String(http.StatusInternalServerError, "db connection failed")
	}
	defer db.Close()
	db.Exec("DELETE FROM booking WHERE id = ?;", booking.ID)
	return c.NoContent(200)
}

func GetBooking(c echo.Context) error {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return c.String(http.StatusInternalServerError, "db connection failed")
	}
	defer db.Close()
	var room models.Room
	c.Bind(room)
	if room.ID == nil {
		return c.String(400, "room id required")
	}
	var booking []struct {
		ID        *int       `json:"id"`
		DateStart *time.Time `json:"date_start"`
		DateEnd   *time.Time `json:"date_end"`
	}
	rows, err := db.Query("SELECT id, datestart, dateend FROM booking WHERE roomid = ? ORDER BY datestart", &room.ID)
	if err != nil {
		return c.String(500, "get rows error")
	}
	defer rows.Close()
	for rows.Next() {
		var book struct {
			ID        *int       `json:"id"`
			DateStart *time.Time `json:"date_start"`
			DateEnd   *time.Time `json:"date_end"`
		}
		err := rows.Scan(&book.ID, &book.DateStart, &book.DateEnd)
		if err != nil {
			return c.String(500, "scan row failed")
		}
		booking = append(booking, book)
	}
	return c.JSON(200, booking)
}
