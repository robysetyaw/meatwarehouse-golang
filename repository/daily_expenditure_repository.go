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
	GetTotalExpenditureByDateRange(startDate time.Time, endDate time.Time) (float64, error)
	GetExpendituresByDateRange(startDate time.Time, endDate time.Time) ([]*model.DailyExpenditureReport, error)
}	

type dailyExpenditureRepository struct {
	db *sql.DB
}

func NewDailyExpenditureRepository(db *sql.DB) DailyExpenditureRepository {
	return &dailyExpenditureRepository{
		db: db,
	}
}

func (repo *dailyExpenditureRepository) GetTotalExpenditureByDateRange(startDate time.Time, endDate time.Time) (float64, error) {
    var total float64
    err := repo.db.QueryRow(`
        SELECT SUM(amount) FROM daily_expenditures
        WHERE DATE(created_at) >= $1 AND DATE(created_at) <= $2
    `, startDate, endDate).Scan(&total)
    if err != nil {
        return 0, fmt.Errorf("failed to get total expenditure: %w", err)
    }
	// fmt.Print(startDate)
    return total, nil
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
		WHERE id = $7 AND is_active = true
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
		WHERE id = $1 AND is_active = true
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
		SELECT id, user_id, amount, description, is_active, created_at, updated_at, created_by, updated_by
		FROM daily_expenditures WHERE is_active = true
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
	UPDATE users
	SET is_active = false
	WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf("failed to delete daily expenditure: %w", err)
	}

	return nil
}

func (repo *dailyExpenditureRepository) GetExpendituresByDateRange(startDate time.Time, endDate time.Time) ([]*model.DailyExpenditureReport, error) {
	var expenditures []*model.DailyExpenditureReport

	rows, err := repo.db.Query(`
		SELECT d.id, d.user_id, u.username, d.amount, d.description, d.created_at, d.updated_at, d.date
		FROM daily_expenditures d
		JOIN users u ON d.user_id = u.id
		WHERE d.date >= $1 AND d.date <= $2 AND d.is_active = true
	`, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to get expenditures by date range: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var expenditure model.DailyExpenditureReport
		if err := rows.Scan(
			&expenditure.ID,
			&expenditure.UserID,
			&expenditure.Username,
			&expenditure.Amount,
			&expenditure.Description,
			&expenditure.CreatedAt,
			&expenditure.UpdatedAt,
			&expenditure.Date,
		); err != nil {
			return nil, fmt.Errorf("failed to scan expenditure row: %w", err)
		}

		expenditures = append(expenditures, &expenditure)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while iterating over expenditure rows: %w", err)
	}

	return expenditures, nil
}

