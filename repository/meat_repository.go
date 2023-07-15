package repository

import (
	"database/sql"
	"time"
	"fmt"
	"errors"
	"enigmacamp.com/final-project/team-4/track-prosto/model"
)

type MeatRepository interface {
	CreateMeat(meat *model.Meat) error
	GetMeatByID(string) (*model.Meat, error)
	GetAllMeats() ([]*model.Meat, error)
	GetMeatByName(string)(*model.Meat, error)
	UpdateMeat(meat *model.Meat) error
	DeleteMeat(string) error
	ReduceStock(meatID string, qty float64) error
	IncreaseStock(meatID string, qty float64) error
}

type meatRepository struct {
	db *sql.DB
}

func NewMeatRepository(db *sql.DB) MeatRepository {
	return &meatRepository{db: db}
}

func (mr *meatRepository) CreateMeat(meat *model.Meat) error {
	query := `
	INSERT INTO meats (id, name, stock, price, is_active, created_at, updated_at, created_by, updated_by)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	createdAt := time.Now()
	updatedAt := time.Now()
	isActive := true
	_, err := mr.db.Exec(query, meat.ID, meat.Name, meat.Stock, meat.Price, isActive, createdAt, updatedAt, meat.CreatedBy, meat.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}

func (r *meatRepository) GetAllMeats() ([]*model.Meat, error) {
	query := `
		SELECT id, name, stock, price, is_active, created_at, updated_at, created_by, updated_by
		FROM meats
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	meats := []*model.Meat{}
	for rows.Next() {
		meat := &model.Meat{}
		err := rows.Scan(
			&meat.ID,
			&meat.Name,
			&meat.Stock,
			&meat.Price,
			&meat.IsActive,
			&meat.CreatedAt,
			&meat.UpdatedAt,
			&meat.CreatedBy,
			&meat.UpdatedBy,
		)
		if err != nil {
			return nil, err
		}
		meats = append(meats, meat)
	}

	return meats, nil
}

func (r *meatRepository) GetMeatByName(name string) (*model.Meat, error) {
	query := `
		SELECT id, name, stock, price, is_active, created_at, updated_at, created_by, updated_by
		FROM meats
		WHERE name = $1 AND is_active = true
	`
	meat := &model.Meat{}
	err := r.db.QueryRow(query, name).Scan(&meat.ID, &meat.Name, &meat.Stock, &meat.Price, &meat.IsActive, &meat.CreatedAt, &meat.UpdatedAt, &meat.CreatedBy, &meat.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Meat not found
		}
		return nil, fmt.Errorf("failed to get meat by meatname: %v", err)
	}
	return meat, nil
}

func (r *meatRepository) GetMeatByID(id string) (*model.Meat, error) {
	query := `
		SELECT id, name, stock, price, is_active, created_at, updated_at, created_by, updated_by
		FROM meats
		WHERE id = $1
	`

	row := r.db.QueryRow(query, id)

	meat := &model.Meat{}
	err := row.Scan(
		&meat.ID,
		&meat.Name,
		&meat.Stock,
		&meat.Price,
		&meat.IsActive,
		&meat.CreatedAt,
		&meat.UpdatedAt,
		&meat.CreatedBy,
		&meat.UpdatedBy,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Return nil if no meat found
		}
		return nil, err
	}

	return meat, nil
}

func (r *meatRepository) DeleteMeat(id string) error {
	query := `
	UPDATE meats
	SET is_active = false
	WHERE id = $1
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *meatRepository) UpdateMeat(meat *model.Meat) error {
	query := `
		UPDATE meats
		SET name = $1, stock = $2, price = $3, is_active = $4, updated_at = $5, updated_by = $6
		WHERE id = $7
	`
	updatedAt := time.Now()

	_, err := r.db.Exec(query, meat.Name, meat.Stock, meat.Price, meat.IsActive, updatedAt, meat.UpdatedBy, meat.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *meatRepository) ReduceStock(meatID string, qty float64) error {
	query := "UPDATE meats SET stock = stock - $1 WHERE id = $2"
	_, err := r.db.Exec(query, qty,meatID)
	if err != nil {
		return err
	}

	return nil
}

func (r *meatRepository) IncreaseStock(meatID string, qty float64) error {
	query := "UPDATE meats SET stock = stock + $1 WHERE id = $2"
	_, err := r.db.Exec(query, qty,meatID)
	if err != nil {
		return err
	}

	return nil
}