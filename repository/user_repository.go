package repository

import (
	"database/sql"
	"log"

	"enigmacamp.com/final-project/team-4/track-prosto/model"
	"enigmacamp.com/final-project/team-4/track-prosto/utils"
)

type UserRepository interface {
	CreateUser(*model.UserModel) error
	UpdateUser(*model.UserModel) error
	DeleteUser(string) error
	GetUserByID(string) (*model.UserModel, error)
	GetAllUsers() ([]*model.UserModel, error)
	GetUserByName(string) (*model.UserModel, error)
}
type ursitoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &ursitoryImpl{
		db: db,
	}
}

func (ur *ursitoryImpl) CreateUser(user *model.UserModel) error {
	query := utils.INSERT_USER

	_, err := ur.db.Exec(query, user.ID, user.Username, user.Password, user.IsActive, user.Role, user.CreatedAt, user.UpdatedAt, user.CreatedBy, user.UpdatedBy)
	if err != nil {
		log.Println("Error CreateUser():", err)
		return err
	}

	return nil
}

func (ur *ursitoryImpl) UpdateUser(user *model.UserModel) error {
	query := utils.UPDATE_USER

	_, err := ur.db.Exec(query, user.ID, user.Username, user.Password, user.IsActive, user.Role, user.UpdatedAt, user.UpdatedBy)
	if err != nil {
		log.Println("Error UpdateUser():", err)
		return err
	}

	return nil
}

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
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		log.Println("Error GetUserByID():", err)
		return nil, err
	}

	return user, nil
}

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
		return nil, err
	}
	defer rows.Close()

	users := []*model.UserModel{}
	for rows.Next() {
		user := &model.UserModel{}
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
