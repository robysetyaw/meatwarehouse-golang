package repository

import (
	"database/sql"
	"fmt"
	"time"

	"enigmacamp.com/final-project/team-4/track-prosto/model"
)

type DailyExpenditureRepository interface {
	CreateDailyExpenditure(expenditure *model.DailyExpenditure) error
	UpdateDailyExpenditure(expenditure *model.DailyExpenditure) error
	GetDailyExpenditureByID(id string) (*model.DailyExpenditure, error)
	GetAllDailyExpenditures() ([]*model.DailyExpenditure, error)
	DeleteDailyExpenditure(id string) error
}

type dailyExpenditureRepository struct {
	db *sql.DB
}

func NewDailyExpenditureRepository(db *sql.DB) DailyExpenditureRepository {
	return &dailyExpenditureRepository{
		db: db,
	}
}

func (repo *dailyExpenditureRepository) CreateDailyExpenditure(expenditure *model.DailyExpenditure) error {
	now := time.Now()
	expenditure.CreatedAt = now
	expenditure.UpdatedAt = now

	// Perform database insert operation
	_, err := repo.db.Exec(`
		INSERT INTO daily_expenditures (id, user_id, amount, description, is_active, created_at, updated_at, created_by, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`, expenditure.ID, expenditure.UserID, expenditure.Amount, expenditure.Description, expenditure.IsActive,  expenditure.CreatedAt, expenditure.UpdatedAt, expenditure.CreatedBy, expenditure.CreatedBy)
	if err != nil {
		return fmt.Errorf("failed to create daily expenditure: %w", err)
	}

	return nil
}

func (repo *dailyExpenditureRepository) UpdateDailyExpenditure(expenditure *model.DailyExpenditure) error {
	// Set updated_at timestamp
	expenditure.UpdatedAt = time.Now()

	// Perform database update operation
	_, err := repo.db.Exec(`
		UPDATE daily_expenditures
		SET amount = $1, description = $2, is_active = $3,  updated_at = $5, updated_by = $6
		WHERE id = $7
	`, expenditure.Amount, expenditure.Description, expenditure.IsActive, expenditure.UpdatedAt, expenditure.UpdatedBy, expenditure.ID)
	if err != nil {
		return fmt.Errorf("failed to update daily expenditure: %w", err)
	}

	return nil
}

func (repo *dailyExpenditureRepository) GetDailyExpenditureByID(id string) (*model.DailyExpenditure, error) {
	var expenditure model.DailyExpenditure

	// Perform database query to retrieve the daily expenditure by ID
	err := repo.db.QueryRow(`
		SELECT id, user_id, amount, description, is_active, role, created_at, updated_at, created_by, updated_by
		FROM daily_expenditures
		WHERE id = $1
	`, id).Scan(
		&expenditure.ID,
		&expenditure.UserID,
		&expenditure.Amount,
		&expenditure.Description,
		&expenditure.IsActive,
		&expenditure.CreatedAt,
		&expenditure.UpdatedAt,
		&expenditure.CreatedBy,
		&expenditure.UpdatedBy,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Daily expenditure not found
		}
		return nil, fmt.Errorf("failed to get daily expenditure by ID: %w", err)
	}

	return &expenditure, nil
}

func (repo *dailyExpenditureRepository) GetAllDailyExpenditures() ([]*model.DailyExpenditure, error) {
	// Perform database query to retrieve all daily expenditures
	rows, err := repo.db.Query(`
		SELECT id, user_id, amount, description, is_active, role, created_at, updated_at, created_by, updated_by
		FROM daily_expenditures
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get all daily expenditures: %w", err)
	}
	defer rows.Close()

	// Iterate over the rows and scan the results into DailyExpenditure objects
	expenditures := make([]*model.DailyExpenditure, 0)
	for rows.Next() {
		var expenditure model.DailyExpenditure
		err := rows.Scan(
			&expenditure.ID,
			&expenditure.UserID,
			&expenditure.Amount,
			&expenditure.Description,
			&expenditure.IsActive,
			&expenditure.CreatedAt,
			&expenditure.UpdatedAt,
			&expenditure.CreatedBy,
			&expenditure.UpdatedBy,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan daily expenditure row: %w", err)
		}
		expenditures = append(expenditures, &expenditure)
	}

	return expenditures, nil
}

func (repo *dailyExpenditureRepository) DeleteDailyExpenditure(id string) error {
	// Perform database delete operation
	_, err := repo.db.Exec(`
		DELETE FROM daily_expenditures WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf("failed to delete daily expenditure: %w", err)
	}

	return nil
}
