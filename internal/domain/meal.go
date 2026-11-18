package domain

import (
	"time"

	"github.com/google/uuid"
)

type MealLog struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserID    uuid.UUID `json:"userId"`
	Text      string    `json:"text"`
	Energy
	Items []MealItem `json:"items"`
}

type MealItem struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Weight int       `json:"weight"`
	Energy
}

type Energy struct {
	KCal         int     `json:"kcal"`
	Fat          float64 `json:"fat"`
	Protein      float64 `json:"protein"`
	Carbohydrate float64 `json:"carbohydrate"`
}

func (m *MealLog) CalculateTotalEnergy() {
	for _, item := range m.Items {
		m.KCal += item.KCal
		m.Carbohydrate += item.Carbohydrate
		m.Fat += item.Fat
		m.Protein += item.Protein
	}
}
