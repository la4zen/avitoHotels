package models

import "time"

type Room struct {
	ID           *int       `json:"id"`
	Discription  *string    `json:"discription"`
	Price        *int       `json:"price"`
	CreationDate *time.Time `json:"creation_date"`
}

type Booking struct {
	ID        *int       `json:"id"`
	RoomID    *int       `json:"room_id"`
	DateStart *time.Time `json:"date_start"`
	DateEnd   *time.Time `json:"date_end"`
}
