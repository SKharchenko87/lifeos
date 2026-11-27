package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"lifeos/internal/domain"

	"github.com/google/uuid"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) Create(ctx context.Context, product *domain.Product) error {
	if product.ID == uuid.Nil {
		return fmt.Errorf("Не заполнен ID")
	}
	query := `INSERT INTO products (id, name, brand, category, barcode, kcal, fat, protein, carbohydrate, source, created_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := p.db.ExecContext(
		ctx,
		query,
		product.ID,
		product.Name,
		product.Brand,
		product.Category,
		product.BarCode,
		product.KCal,
		product.Fat,
		product.Protein,
		product.Carbohydrate,
		product.Source,
		product.CreatedAt,
	)
	return err
}

func (p *ProductRepository) SearchBestMatch(ctx context.Context, text string) (*domain.Product, error) {
	query := `select 
		id, 
		name, 
		COALESCE(brand,''), 
		COALESCE(category,''), 
		COALESCE(barcode,''), 
		kcal, 
		fat, 
		protein, 
		carbohydrate, 
		source, 
		created_at 
	from products 
	where name % $1 
	order by 
	similarity(name, $1) desc 
	limit 1`

	res := &domain.Product{}

	err := p.db.QueryRowContext(ctx, query, text).Scan(
		&res.ID,
		&res.Name,
		&res.Brand,
		&res.Category,
		&res.BarCode,
		&res.KCal,
		&res.Fat,
		&res.Protein,
		&res.Carbohydrate,
		&res.Source,
		&res.CreatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return res, nil
}
