package handlers

import (
	"database/sql"
	"encoding/json"
	"event-booking-api/internal/models"
	"net/http"
)

type EventHandler struct {
	DB *sql.DB
}

func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var e models.Event
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	query := `INSERT INTO events (title, description, date, location, available_seats)
              VALUES ($1, $2, $3, $4,  $5) RETURNING id, created_at`

	err := h.DB.QueryRow(query, e.Title, e.Description, e.Date, e.Location, e.AvailableSeats).Scan(&e.ID, &e.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&e)
}

func (h *EventHandler) GetEvents(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query(`SELECT id, title, description, date, location, available_seats, created_at FROM events`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var e models.Event
		if err := rows.Scan(&e.ID, &e.Title, &e.Date, &e.Location, &e.AvailableSeats, &e.CreatedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		events = append(events, e)
	}
	json.NewEncoder(w).Encode(events)
}
