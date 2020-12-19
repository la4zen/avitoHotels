package models

import "time"

type Room struct {
	ID          int       `json:"id,omitempty"`
	Discription string    `json:"discription,omitempty"`
	Price       int       `json:"price,omitempty"`
	CreatinDate time.Time `json:"creatin_date,omitempty"`
}

type Booking struct {
	ID        int       `json:"id,omitempty"`
	RoomId    int       `json:"room_id,omitempty"`
	DateStart time.Time `json:"date_start,omitempty"`
	DateEnd   time.Time `json:"date_end,omitempty"`
}
