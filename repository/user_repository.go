package repository

import (
	"database/sql"
	"log"
	

	"enigmacamp.com/final-project/team-4/track-prosto/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) CreateUser(user *model.User) error {
	query := `
		INSERT INTO users (id, username, password, is_active, role, created_at, updated_at, created_by, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := repo.db.Exec(query, user.ID, user.Username, user.Password, user.IsActive, user.Role, user.CreatedAt, user.UpdatedAt, user.CreatedBy, user.UpdatedBy)
	if err != nil {
		log.Println("Error creating user:", err)
		return err
	}

	return nil
}

func (repo *UserRepository) UpdateUser(user *model.User) error {
	query := `
		UPDATE users
		SET username = $2, password = $3, is_active = $4, role = $5, updated_at = $6, updated_by = $7
		WHERE id = $1
	`

	_, err := repo.db.Exec(query, user.ID, user.Username, user.Password, user.IsActive, user.Role, user.UpdatedAt, user.UpdatedBy)
	if err != nil {
		log.Println("Error updating user:", err)
		return err
	}

	return nil
}

func (repo *UserRepository) DeleteUser(userID string) error {
	query := `
		DELETE FROM users WHERE id = $1
	`

	_, err := repo.db.Exec(query, userID)
	if err != nil {
		log.Println("Error deleting user:", err)
		return err
	}

	return nil
}

func (repo *UserRepository) GetUserByID(userID string) (*model.User, error) {
	query := `
		SELECT id, username, password, is_active, role, created_at, updated_at, created_by, updated_by
		FROM users
		WHERE id = $1
	`

	row := repo.db.QueryRow(query, userID)

	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.IsActive, &user.Role, &user.CreatedAt, &user.UpdatedAt, &user.CreatedBy, &user.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		log.Println("Error retrieving user:", err)
		return nil, err
	}

	return user, nil
}

func (repo *UserRepository) GetAllUsers() ([]*model.User, error) {
	query := `
		SELECT id, username, password, is_active, role, created_at, updated_at, created_by, updated_by
		FROM users
	`

	rows, err := repo.db.Query(query)
	if err != nil {
		log.Println("Error retrieving users:", err)
		return nil, err
	}
	defer rows.Close()

	users := []*model.User{}
	for rows.Next() {
		user := &model.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.IsActive, &user.Role, &user.CreatedAt, &user.UpdatedAt, &user.CreatedBy, &user.UpdatedBy)
		if err != nil {
			log.Println("Error scanning user row:", err)
			continue
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating over user rows:", err)
		return nil, err
	}

	return users, nil
}
