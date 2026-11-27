package domain

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Brand    string    `json:"brand"`
	Category string    `json:"category"`
	BarCode  string    `json:"bar_code"`
	Energy
	Source    string    `json:"source"`
	CreatedAt time.Time `json:"created_at"`
}
