package repository

import (
	"database/sql"
	"log"

	"enigmacamp.com/final-project/team-4/track-prosto/model"
)

type DailyExpenditureRepository struct {
	db *sql.DB
}

func NewDailyExpenditureRepository(db *sql.DB) *DailyExpenditureRepository {
	return &DailyExpenditureRepository{db: db}
}

func (repo *DailyExpenditureRepository) CreateDailyExpenditure(dailyExpenditure *model.DailyExpenditure) error {
	query := `
		INSERT INTO daily_expenditures (id, user_id, amount, description, is_active, created_at, updated_at, created_by, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := repo.db.Exec(query, dailyExpenditure.ID, dailyExpenditure.UserID, dailyExpenditure.Amount, dailyExpenditure.Description, dailyExpenditure.IsActive, dailyExpenditure.CreatedAt, dailyExpenditure.UpdatedAt, dailyExpenditure.CreatedBy, dailyExpenditure.UpdatedBy)
	if err != nil {
		log.Println("Error creating daily expenditure:", err)
		return err
	}

	return nil
}

func (repo *DailyExpenditureRepository) UpdateDailyExpenditure(dailyExpenditure *model.DailyExpenditure) error {
	query := `
		UPDATE daily_expenditures
		SET user_id = $2, amount = $3, description = $4, is_active = $5, updated_at = $6, updated_by = $7
		WHERE id = $1
	`

	_, err := repo.db.Exec(query, dailyExpenditure.ID, dailyExpenditure.UserID, dailyExpenditure.Amount, dailyExpenditure.Description, dailyExpenditure.IsActive, dailyExpenditure.UpdatedAt, dailyExpenditure.UpdatedBy)
	if err != nil {
		log.Println("Error updating daily expenditure:", err)
		return err
	}

	return nil
}

func (repo *DailyExpenditureRepository) DeleteDailyExpenditure(dailyExpenditureID string) error {
	query := `
		DELETE FROM daily_expenditures WHERE id = $1
	`

	_, err := repo.db.Exec(query, dailyExpenditureID)
	if err != nil {
		log.Println("Error deleting daily expenditure:", err)
		return err
	}

	return nil
}

func (repo *DailyExpenditureRepository) GetDailyExpenditureByID(dailyExpenditureID string) (*model.DailyExpenditure, error) {
	query := `
		SELECT id, user_id, amount, description, is_active, created_at, updated_at, created_by, updated_by
		FROM daily_expenditures
		WHERE id = $1
	`

	row := repo.db.QueryRow(query, dailyExpenditureID)

	dailyExpenditure := &model.DailyExpenditure{}
	err := row.Scan(&dailyExpenditure.ID, &dailyExpenditure.UserID, &dailyExpenditure.Amount, &dailyExpenditure.Description, &dailyExpenditure.IsActive, &dailyExpenditure.CreatedAt, &dailyExpenditure.UpdatedAt, &dailyExpenditure.CreatedBy, &dailyExpenditure.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // DailyExpenditure not found
		}
		log.Println("Error retrieving daily expenditure:", err)
		return nil, err
	}

	return dailyExpenditure, nil
}

func (repo *DailyExpenditureRepository) GetAllDailyExpenditures() ([]*model.DailyExpenditure, error) {
	query := `
		SELECT id, user_id, amount, description, is_active, created_at, updated_at, created_by, updated_by
		FROM daily_expenditures
	`

	rows, err := repo.db.Query(query)
	if err != nil {
		log.Println("Error retrieving daily expenditures:", err)
		return nil, err
	}
	defer rows.Close()

	dailyExpenditures := []*model.DailyExpenditure{}
	for rows.Next() {
		dailyExpenditure := &model.DailyExpenditure{}
		err := rows.Scan(&dailyExpenditure.ID, &dailyExpenditure.UserID, &dailyExpenditure.Amount, &dailyExpenditure.Description, &dailyExpenditure.IsActive, &dailyExpenditure.CreatedAt, &dailyExpenditure.UpdatedAt, &dailyExpenditure.CreatedBy, &dailyExpenditure.UpdatedBy)
		if err != nil {
			log.Println("Error scanning daily expenditure row:", err)
			continue
		}
		dailyExpenditures = append(dailyExpenditures, dailyExpenditure)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating over daily expenditure rows:", err)
		return nil, err
	}

	return dailyExpenditures, nil
}
