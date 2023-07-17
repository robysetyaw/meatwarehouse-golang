package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"time"

	"enigmacamp.com/final-project/team-4/track-prosto/model"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	GetUserByID(id string) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
	DeleteUser(id string) error
	GetByUsername(username string) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *model.User) error {
	query := `
		INSERT INTO users (id, username, password, is_active, role, created_at, updated_at, created_by, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	
	_, err := r.db.Exec(query, user.ID, user.Username, user.Password, user.IsActive, user.Role, user.CreatedAt, user.CreatedAt, user.CreatedBy, user.CreatedBy)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) UpdateUser(user *model.User) error {
	query := `
		UPDATE users
		SET username = $1, password = $2, is_active = $3, role = $4, updated_at = $5, updated_by = $6
		WHERE id = $7 AND is_active = true
	`

	updatedAt := time.Now()

	_, err := r.db.Exec(query, user.Username, user.Password, user.IsActive, user.Role, updatedAt, user.UpdatedBy, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetUserByID(id string) (*model.User, error) {
	query := `
		SELECT id, username, password, is_active, role, created_at, updated_at, created_by, updated_by
		FROM users
		WHERE id = $1
	`

	row := r.db.QueryRow(query, id)

	user := &model.User{}
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.IsActive,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.CreatedBy,
		&user.UpdatedBy,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Return nil if no user found
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetByUsername(username string) (*model.User, error) {
	query := `
		SELECT id, username, password, is_active, role, created_at, updated_at, created_by, updated_by
		FROM users
		WHERE username = $1 AND is_active = true
	`
	user := &model.User{}
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.IsActive, &user.Role, &user.CreatedAt, &user.UpdatedAt, &user.CreatedBy, &user.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, fmt.Errorf("failed to get user by username: %v", err)
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]*model.User, error) {
	query := `
		SELECT id, username, password, is_active, role, created_at, updated_at, created_by, updated_by
		FROM users WHERE is_active = true
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*model.User{}
	for rows.Next() {
		user := &model.User{}
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.IsActive,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.CreatedBy,
			&user.UpdatedBy,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) DeleteUser(username string) error {
	query := `
	UPDATE users
	SET is_active = false
	WHERE username = $1
	`

	_, err := r.db.Exec(query, username)
	if err != nil {
		return err
	}

	return nil
}
