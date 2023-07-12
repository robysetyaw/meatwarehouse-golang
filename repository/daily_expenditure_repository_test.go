package repository

import (
	"errors"
	"testing"
	"time"

	"enigmacamp.com/final-project/team-4/track-prosto/model"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestDailyExpenditureRepository_CreateDailyExpenditure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	defer db.Close()

	repo := NewDailyExpenditureRepository(db)

	// Mock data
	dailyExpenditure := &model.DailyExpenditure{
		ID:          "1",
		UserID:      "user1",
		Amount:      100.50,
		Description: "Test expenditure",
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CreatedBy:   "admin",
		UpdatedBy:   "admin",
	}

	// Expectation
	mock.ExpectExec("INSERT INTO daily_expenditures").
		WithArgs(dailyExpenditure.ID, dailyExpenditure.UserID, dailyExpenditure.Amount, dailyExpenditure.Description, dailyExpenditure.IsActive, dailyExpenditure.CreatedAt, dailyExpenditure.UpdatedAt, dailyExpenditure.CreatedBy, dailyExpenditure.UpdatedBy).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the actual method
	err = repo.CreateDailyExpenditure(dailyExpenditure)
	assert.NoError(t, err)

	// Verify the expectations
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestDailyExpenditureRepository_CreateDailyExpenditure_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	defer db.Close()

	repo := NewDailyExpenditureRepository(db)

	// Mock data
	dailyExpenditure := &model.DailyExpenditure{
		ID:          "1",
		UserID:      "user1",
		Amount:      100.50,
		Description: "Test expenditure",
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CreatedBy:   "admin",
		UpdatedBy:   "admin",
	}

	// Expectation
	mock.ExpectExec("INSERT INTO daily_expenditures").
		WithArgs(dailyExpenditure.ID, dailyExpenditure.UserID, dailyExpenditure.Amount, dailyExpenditure.Description, dailyExpenditure.IsActive, dailyExpenditure.CreatedAt, dailyExpenditure.UpdatedAt, dailyExpenditure.CreatedBy, dailyExpenditure.UpdatedBy).
		WillReturnError(errors.New("database error"))

	// Call the actual method
	err = repo.CreateDailyExpenditure(dailyExpenditure)
	assert.Error(t, err)
	assert.EqualError(t, err, "database error")

	// Verify the expectations
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestDailyExpenditureRepository_UpdateDailyExpenditure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	defer db.Close()

	repo := NewDailyExpenditureRepository(db)

	// Mock data
	dailyExpenditure := &model.DailyExpenditure{
		ID:          "1",
		UserID:      "user1",
		Amount:      200.75,
		Description: "Updated expenditure",
		IsActive:    true,
		UpdatedAt:   time.Now(),
		UpdatedBy:   "admin",
	}

	// Expectation
	mock.ExpectExec("UPDATE daily_expenditures").
		WithArgs(dailyExpenditure.ID, dailyExpenditure.UserID, dailyExpenditure.Amount, dailyExpenditure.Description, dailyExpenditure.IsActive, dailyExpenditure.UpdatedAt, dailyExpenditure.UpdatedBy).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the actual method
	err = repo.UpdateDailyExpenditure(dailyExpenditure)
	assert.NoError(t, err)

	// Verify the expectations
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestDailyExpenditureRepository_UpdateDailyExpenditure_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	defer db.Close()

	repo := NewDailyExpenditureRepository(db)

	// Mock data
	dailyExpenditure := &model.DailyExpenditure{
		ID:          "1",
		UserID:      "user1",
		Amount:      200.75,
		Description: "Updated expenditure",
		IsActive:    true,
		UpdatedAt:   time.Now(),
		UpdatedBy:   "admin",
	}

	// Expectation
	mock.ExpectExec("UPDATE daily_expenditures").
		WithArgs(dailyExpenditure.ID, dailyExpenditure.UserID, dailyExpenditure.Amount, dailyExpenditure.Description, dailyExpenditure.IsActive, dailyExpenditure.UpdatedAt, dailyExpenditure.UpdatedBy).
		WillReturnError(errors.New("database error"))

	// Call the actual method
	err = repo.UpdateDailyExpenditure(dailyExpenditure)
	assert.Error(t, err)
	assert.EqualError(t, err, "database error")

	// Verify the expectations
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
func TestDailyExpenditureRepository_DeleteDailyExpenditure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	defer db.Close()

	repo := NewDailyExpenditureRepository(db)

	// Mock data
	dailyExpenditureID := "1"

	// Expectation
	mock.ExpectExec("DELETE FROM daily_expenditures").
		WithArgs(dailyExpenditureID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the actual method
	err = repo.DeleteDailyExpenditure(dailyExpenditureID)
	assert.NoError(t, err)

	// Verify the expectations
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestDailyExpenditureRepository_DeleteDailyExpenditure_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	defer db.Close()

	repo := NewDailyExpenditureRepository(db)

	// Mock data
	dailyExpenditureID := "1"

	// Expectation
	mock.ExpectExec("DELETE FROM daily_expenditures").
		WithArgs(dailyExpenditureID).
		WillReturnError(errors.New("database error"))

	// Call the actual method
	err = repo.DeleteDailyExpenditure(dailyExpenditureID)
	assert.Error(t, err)
	assert.EqualError(t, err, "database error")

	// Verify the expectations
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
func TestDailyExpenditureRepository_GetDailyExpenditureByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	defer db.Close()

	repo := NewDailyExpenditureRepository(db)

	// Mock data
	dailyExpenditureID := "1"
	columns := []string{"id", "user_id", "amount", "description", "is_active", "created_at", "updated_at", "created_by", "updated_by"}
	rows := sqlmock.NewRows(columns).
		AddRow(dailyExpenditureID, "user1", 150.25, "Test expenditure", true, time.Now(), time.Now(), "admin", "admin")

	// Expectation
	mock.ExpectQuery("SELECT (.+) FROM daily_expenditures").
		WithArgs(dailyExpenditureID).
		WillReturnRows(rows)

	// Call the actual method
	dailyExpenditure, err := repo.GetDailyExpenditureByID(dailyExpenditureID)
	assert.NoError(t, err)
	assert.NotNil(t, dailyExpenditure)
	assert.Equal(t, dailyExpenditureID, dailyExpenditure.ID)

	// Verify the expectations
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestDailyExpenditureRepository_GetDailyExpenditureByID_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	defer db.Close()

	repo := NewDailyExpenditureRepository(db)

	// Mock data
	dailyExpenditureID := "1"

	// Expectation
	mock.ExpectQuery("SELECT (.+) FROM daily_expenditures").
		WithArgs(dailyExpenditureID).
		WillReturnError(errors.New("database error"))

	// Call the actual method
	dailyExpenditure, err := repo.GetDailyExpenditureByID(dailyExpenditureID)
	assert.Error(t, err)
	assert.Nil(t, dailyExpenditure)
	assert.EqualError(t, err, "database error")

	// Verify the expectations
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestDailyExpenditureRepository_GetAllDailyExpenditures(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	defer db.Close()

	repo := NewDailyExpenditureRepository(db)

	// Mock data
	columns := []string{"id", "user_id", "amount", "description", "is_active", "created_at", "updated_at", "created_by", "updated_by"}
	rows := sqlmock.NewRows(columns).
		AddRow("1", "user1", 150.25, "Expenditure 1", true, time.Now(), time.Now(), "admin", "admin").
		AddRow("2", "user2", 200.50, "Expenditure 2", true, time.Now(), time.Now(), "admin", "admin")

	// Expectation
	mock.ExpectQuery("SELECT (.+) FROM daily_expenditures").
		WillReturnRows(rows)

	// Call the actual method
	dailyExpenditures, err := repo.GetAllDailyExpenditures()
	assert.NoError(t, err)
	assert.NotNil(t, dailyExpenditures)
	assert.Len(t, dailyExpenditures, 2)

	// Verify the expectations
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestDailyExpenditureRepository_GetAllDailyExpenditures_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	defer db.Close()

	repo := NewDailyExpenditureRepository(db)

	// Expectation
	mock.ExpectQuery("SELECT (.+) FROM daily_expenditures").
		WillReturnError(errors.New("database error"))

	// Call the actual method
	dailyExpenditures, err := repo.GetAllDailyExpenditures()
	assert.Error(t, err)
	assert.Nil(t, dailyExpenditures)
	assert.EqualError(t, err, "database error")

	// Verify the expectations
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
