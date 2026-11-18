package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"lifeos/internal/domain"

	"github.com/google/uuid"
)

type MealRepository struct {
	db *sql.DB
}

func NewMealRepository(db *sql.DB) *MealRepository {
	return &MealRepository{db: db}
}

func (r *MealRepository) Create(ctx context.Context, meal *domain.MealLog) error {
	var err error

	if meal.ID == uuid.Nil {
		return fmt.Errorf("Не заполнен ID")
	}

	items, err := json.Marshal(meal.Items)
	if err != nil {
		return err
	}

	query := "insert into meal_logs (id, user_id, text, kcal, fat, protein, carbohydrate, items) values ($1, $2, $3, $4, $5, $6, $7, $8)"
	_, err = r.db.ExecContext(ctx, query, meal.ID, meal.UserID, meal.Text, meal.KCal, meal.Fat, meal.Protein, meal.Carbohydrate, items)
	if err != nil {
		return err
	}
	return nil
}
