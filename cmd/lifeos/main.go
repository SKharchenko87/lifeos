package main

import (
	"context"
	"database/sql"
	"fmt"
	"lifeos/internal/domain"
	"lifeos/internal/repository"
	"log"
	"time"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	var err error
	ctx := context.Background()
	Client, err := sql.Open("pgx", "postgres://lifeos:lifeos@localhost:15432/lifeos?sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	if err := Client.PingContext(ctx); err != nil {
		log.Fatalf("ping DB failed: %v", err)
	}
	defer Client.Close()

	rep := repository.NewProductRepository(Client)
	product0 := domain.Product{
		ID:           uuid.New(),
		Name:         "Творог 5%",
		Category:     "Молочные продукты",
		Brand:        "",
		BarCode:      "",
		KCal:         121,
		Protein:      16.0,
		Fat:          5,
		Carbohydrate: 3,
		Source:       "Ручной ввод",
		CreatedAt:    time.Now(),
	}
	product1 := domain.Product{
		ID:           uuid.New(),
		Name:         "Филе куриной грудкии",
		Category:     "Мясо",
		Brand:        "",
		BarCode:      "",
		KCal:         113,
		Protein:      23.46,
		Fat:          1.9,
		Carbohydrate: 0.4,
		Source:       "Ручной ввод",
		CreatedAt:    time.Now(),
	}

	err = rep.Create(ctx, &product0)
	if err != nil {
		log.Fatalf("Ошибка при записи продукта 0: %v", err)
	}

	err = rep.Create(ctx, &product1)
	if err != nil {
		log.Fatalf("Ошибка при записи продукта 1: %v", err)
	}

	p, err := rep.SearchBestMatch(ctx, "Творожок")
	if err != nil {
		log.Fatalf("Ошибка поиска: %v", err)
	}
	fmt.Printf("%v", p)

	// rep := repository.NewMealRepository(Client)
	// items := make([]domain.MealItem, 3)
	// items[0].KCal = 100
	// items[0].Fat = 1.0
	// items[0].Protein = 2.0
	// items[0].Carbohydrate = 3.0
	// items[0].Weight = 50

	// items[1].KCal = 101
	// items[1].Fat = 1.1
	// items[1].Protein = 2.1
	// items[1].Carbohydrate = 3.1
	// items[1].Weight = 51

	// items[2].KCal = 102
	// items[2].Fat = 1.2
	// items[2].Protein = 2.2
	// items[2].Carbohydrate = 3.2
	// items[2].Weight = 52

	// v := domain.MealLog{ID: uuid.New(), UserID: uuid.New(), Text: "Творог %5 100 грам", Items: items}
	// v.CalculateTotalEnergy()

	// err = rep.Create(ctx, &v)
	// if err != nil {
	// 	log.Fatalf("create failed: %v", err)
	// }

}
