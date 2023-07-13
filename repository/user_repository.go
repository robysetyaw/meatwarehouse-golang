package repository

import (
	"database/sql"
<<<<<<< HEAD
	"log"
=======
	"errors"
	"fmt"

	"time"
>>>>>>> origin/dev-borr

	"enigmacamp.com/final-project/team-4/track-prosto/model"
	"enigmacamp.com/final-project/team-4/track-prosto/utils"
)

type UserRepository interface {
<<<<<<< HEAD
	CreateUser(*model.UserModel) error
	UpdateUser(*model.UserModel) error
	DeleteUser(string) error
	GetUserByID(string) (*model.UserModel, error)
	GetAllUsers() ([]*model.UserModel, error)
	GetUserByName(string) (*model.UserModel, error)
}
type ursitoryImpl struct {
=======
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	GetUserByID(id string) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
	DeleteUser(id string) error
	GetByUsername(username string) (*model.User, error)

}

type userRepository struct {
>>>>>>> origin/dev-borr
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
<<<<<<< HEAD
	return &ursitoryImpl{
		db: db,
	}
}

func (ur *ursitoryImpl) CreateUser(user *model.UserModel) error {
	query := utils.INSERT_USER

	_, err := ur.db.Exec(query, user.ID, user.Username, user.Password, user.IsActive, user.Role, user.CreatedAt, user.UpdatedAt, user.CreatedBy, user.UpdatedBy)
	if err != nil {
		log.Println("Error CreateUser():", err)
=======
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *model.User) error {
	query := `
		INSERT INTO users (id, username, password, is_active, role, created_at, updated_at, created_by, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := r.db.Exec(query, user.ID, user.Username, user.Password, user.IsActive, user.Role, user.CreatedAt, user.CreatedAt, user.CreatedBy, user.CreatedBy)
	if err != nil {
>>>>>>> origin/dev-borr
		return err
	}

	return nil
}

<<<<<<< HEAD
func (ur *ursitoryImpl) UpdateUser(user *model.UserModel) error {
	query := utils.UPDATE_USER

	_, err := ur.db.Exec(query, user.ID, user.Username, user.Password, user.IsActive, user.Role, user.UpdatedAt, user.UpdatedBy)
	if err != nil {
		log.Println("Error UpdateUser():", err)
=======
func (r *userRepository) UpdateUser(user *model.User) error {
	query := `
		UPDATE users
		SET username = $1, password = $2, is_active = $3, role = $4, updated_at = $5, updated_by = $6
		WHERE id = $7
	`

	updatedAt := time.Now()

	_, err := r.db.Exec(query, user.Username, user.Password, user.IsActive, user.Role, updatedAt, user.UpdatedBy, user.ID)
	if err != nil {
>>>>>>> origin/dev-borr
		return err
	}

	return nil
}

<<<<<<< HEAD
func (ur *ursitoryImpl) DeleteUser(userID string) error {
	query := utils.DELETE_USER

	_, err := ur.db.Exec(query, userID)
	if err != nil {
		log.Println("Error DeleteUser():", err)
		return err
	}

	return nil
}

func (ur *ursitoryImpl) GetUserByID(userID string) (*model.UserModel, error) {
	query := utils.GET_USER_BY_ID

	row := ur.db.QueryRow(query, userID)

	user := &model.UserModel{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.IsActive, &user.Role, &user.CreatedAt, &user.UpdatedAt, &user.CreatedBy, &user.UpdatedBy)
=======
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
>>>>>>> origin/dev-borr
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Return nil if no user found
		}
<<<<<<< HEAD
		log.Println("Error GetUserByID():", err)
=======
>>>>>>> origin/dev-borr
		return nil, err
	}

	return user, nil
}

<<<<<<< HEAD
func (ur *ursitoryImpl) GetUserByName(userName string) (*model.UserModel, error) {
	query := utils.GET_USER_BY_NAME

	row := ur.db.QueryRow(query, userName)

	user := &model.UserModel{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.IsActive, &user.Role, &user.CreatedAt, &user.UpdatedAt, &user.CreatedBy, &user.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		log.Println("Error GetUserByName():", err)
		return nil, err
	}

	return user, nil
}

func (ur *ursitoryImpl) GetAllUsers() ([]*model.UserModel, error) {
	query := utils.GET_ALL_USER

	rows, err := ur.db.Query(query)
	if err != nil {
		log.Println("Error GetAllUsers():", err)
=======
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
		FROM users
	`

	rows, err := r.db.Query(query)
	if err != nil {
>>>>>>> origin/dev-borr
		return nil, err
	}
	defer rows.Close()

	users := []*model.UserModel{}
	for rows.Next() {
<<<<<<< HEAD
		user := &model.UserModel{}
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.IsActive, &user.Role, &user.CreatedAt, &user.UpdatedAt, &user.CreatedBy, &user.UpdatedBy)
=======
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
>>>>>>> origin/dev-borr
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) DeleteUser(id string) error {
	query := `
	UPDATE users
	SET is_active = false
	WHERE id = $1
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
