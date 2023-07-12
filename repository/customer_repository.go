package repository

import "database/sql"

type CustomerRepository interface {
}

type customerRepositoryImpl struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepositoryImpl{
		db: db,
	}
}
