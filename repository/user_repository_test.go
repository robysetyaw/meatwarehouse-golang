package repository

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"enigmacamp.com/final-project/team-4/track-prosto/model"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_CreateUser_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	// Mock data
	user := &model.User{
		ID:        "1",
		Username:  "john_doe",
		Password:  "password",
		IsActive:  true,
		Role:      "admin",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: "admin",
		UpdatedBy: "admin",
	}

	// Expectation
	mock.ExpectExec("INSERT INTO users").
		WithArgs(user.ID, user.Username, user.Password, user.IsActive, user.Role, user.CreatedAt, user.UpdatedAt, user.CreatedBy, user.UpdatedBy).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Actual test
	err = repo.CreateUser(user)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestUserRepository_CreateUser_Failure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	// Mock data
	user := &model.User{
		ID:        "1",
		Username:  "john_doe",
		Password:  "password",
		IsActive:  true,
		Role:      "admin",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: "admin",
		UpdatedBy: "admin",
	}

	// Expectation
	mock.ExpectExec("INSERT INTO users").
		WithArgs(user.ID, user.Username, user.Password, user.IsActive, user.Role, user.CreatedAt, user.UpdatedAt, user.CreatedBy, user.UpdatedBy).
		WillReturnError(errors.New("database error"))

	// Actual test
	err = repo.CreateUser(user)
	assert.Error(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestUserRepository_GetUserByID_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	// Mock data
	userID := "1"
	columns := []string{"id", "username", "password", "is_active", "role", "created_at", "updated_at", "created_by", "updated_by"}
	rows := sqlmock.NewRows(columns).
		AddRow(userID, "john_doe", "password", true, "admin", time.Now(), time.Now(), "admin", "admin")

	// Expectation
	mock.ExpectQuery("SELECT (.+) FROM users").
		WithArgs(userID).
		WillReturnRows(rows)

	// Actual test
	result, err := repo.GetUserByID(userID)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestUserRepository_GetUserByID_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	// Mock data
	userID := "1"

	// Expectation
	mock.ExpectQuery("SELECT (.+) FROM users").
		WithArgs(userID).
		WillReturnError(sql.ErrNoRows)

	// Actual test
	result, err := repo.GetUserByID(userID)
	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestUserRepository_UpdateUser_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	// Mock data
	user := &model.User{
		ID:        "1",
		Username:  "john_doe",
		Password:  "password",
		IsActive:  true,
		Role:      "admin",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: "admin",
		UpdatedBy: "admin",
	}

	// Expectation
	mock.ExpectExec("UPDATE users").
		WithArgs(user.ID, user.Username, user.Password, user.IsActive, user.Role, user.UpdatedAt, user.UpdatedBy).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Actual test
	err = repo.UpdateUser(user)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}



func TestUserRepository_UpdateUser_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	// Mock data
	user := &model.User{
		ID:        "1",
		Username:  "john_doe",
		Password:  "password",
		IsActive:  true,
		Role:      "admin",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: "admin",
		UpdatedBy: "admin",
	}

	// Expectation
	mock.ExpectExec("UPDATE users").
		WithArgs(user.ID, user.Username, user.Password, user.IsActive, user.Role, user.UpdatedAt, user.UpdatedBy).
		WillReturnError(errors.New("database error"))

	// Call the actual method
	err = repo.UpdateUser(user)
	assert.Error(t, err)
	assert.EqualError(t, err, "database error")

	// Verify the expectations
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}






func TestUserRepository_DeleteUser_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	// Mock data
	userID := "1"

	// Expectation
	mock.ExpectExec("DELETE FROM users").
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Actual test
	err = repo.DeleteUser(userID)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestUserRepository_DeleteUser_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	// Mock data
	userID := "1"

	// Expectation
	mock.ExpectExec("DELETE FROM users").
		WithArgs(userID).
		WillReturnError(errors.New("database error"))

	// Call the actual method
	err = repo.DeleteUser(userID)
	assert.Error(t, err)
	assert.EqualError(t, err, "database error")

	// Verify the expectations
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}


func TestUserRepository_GetAllUsers_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	// Mock data
	columns := []string{"id", "username", "password", "is_active", "role", "created_at", "updated_at", "created_by", "updated_by"}
	rows := sqlmock.NewRows(columns).
		AddRow("1", "john_doe", "password", true, "admin", time.Now(), time.Now(), "admin", "admin").
		AddRow("2", "jane_doe", "password", true, "user", time.Now(), time.Now(), "admin", "admin")

	// Expectation
	mock.ExpectQuery("SELECT (.+) FROM users").
		WillReturnRows(rows)

	// Actual test
	results, err := repo.GetAllUsers()
	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.Len(t, results, 2)
}

func TestUserRepository_GetAllUsers_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	// Expectation
	mock.ExpectQuery("SELECT (.+) FROM users").
		WillReturnError(errors.New("database error"))

	// Call the actual method
	users, err := repo.GetAllUsers()
	assert.Error(t, err)
	assert.Nil(t, users)
	assert.EqualError(t, err, "database error")

	// Verify the expectations
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}


// ...

// Add more unit tests for other repository functions

// ...
