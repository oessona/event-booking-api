package models

import "time"

type Event struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Date          time.Time `json:"date"`
	Location      string    `json:"location"`
	AvailableSeats int      `json:"available_seats"`
	CreatedAt     time.Time `json:"created_at"`
}
