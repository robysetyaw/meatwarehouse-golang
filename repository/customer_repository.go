package repository

import (
	"database/sql"
	"fmt"

	"enigmacamp.com/final-project/team-4/track-prosto/model"
)

type CustomerRepository interface {
	CreateCustomer(*model.CustomerModel) error
	UpdateCustomer(*model.CustomerModel) error
	GetCustomerById(string) (*model.CustomerModel, error)
	GetCustomerByName(string) (*model.CustomerModel, error)
	GetAllCustomer() ([]*model.CustomerModel, error)
	DeleteCustomer(string) error
}

type customerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}

func (repo *customerRepository) CreateCustomer(customer *model.CustomerModel) error {
	// Perform database insert operation
	_, err := repo.db.Exec(`
		INSERT INTO customers (id, fullname, address, company_id, phone_number, created_at, updated_at, created_by, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`, customer.Id, customer.FullName, customer.Address, customer.CompanyId, customer.PhoneNumber, customer.CreatedAt, customer.UpdatedAt, customer.CreatedBy, customer.CreatedBy)
	if err != nil {
		return fmt.Errorf("failed to create customer: %w", err)
	}

	return nil
}

func (repo *customerRepository) UpdateCustomer(customer *model.CustomerModel) error {
	// Perform database update operation
	_, err := repo.db.Exec(`
		UPDATE customers
		SET fullname = $1, address = $2, phone_number = $3,  updated_at = $4, updated_by = $5
		WHERE id = $6
	`, customer.FullName, customer.Address, customer.PhoneNumber, customer.UpdatedAt, customer.UpdatedBy, customer.Id)
	if err != nil {
		return fmt.Errorf("failed to update customer: %w", err)
	}

	return nil
}

func (repo *customerRepository) GetCustomerById(id string) (*model.CustomerModel, error) {
	var customer model.CustomerModel

	// Perform database query to retrieve the customer by ID
	err := repo.db.QueryRow(`
		SELECT id, fullname, address, company_id, phone_number, created_at, updated_at, created_by, updated_by
		FROM customers
		WHERE id = $1
	`, id).Scan(
		&customer.Id,
		&customer.FullName,
		&customer.Address,
		&customer.CompanyId,
		&customer.PhoneNumber,
		&customer.CreatedAt,
		&customer.UpdatedAt,
		&customer.CreatedBy,
		&customer.UpdatedBy,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // customer not found
		}
		return nil, fmt.Errorf("failed to get customer by ID: %w", err)
	}

	return &customer, nil
}

func (repo *customerRepository) GetCustomerByName(customerName string) (*model.CustomerModel, error) {
	var customer model.CustomerModel

	// Perform database query to retrieve the customer by name
	err := repo.db.QueryRow(`
		SELECT id, fullname, address, company_id, phone_number, created_at, updated_at, created_by, updated_by
		FROM customers
		WHERE fullname = $1
	`, customerName).Scan(
		&customer.Id,
		&customer.FullName,
		&customer.Address,
		&customer.CompanyId,
		&customer.PhoneNumber,
		&customer.CreatedAt,
		&customer.UpdatedAt,
		&customer.CreatedBy,
		&customer.UpdatedBy,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("notfoundcustomer %w", err)
		}
		return nil, fmt.Errorf("failed to get customer by name: %w", err)
	}

	return &customer, nil
}

func (repo *customerRepository) GetAllCustomer() ([]*model.CustomerModel, error) {
	// Perform database query to retrieve all customers
	rows, err := repo.db.Query(`
	SELECT id, fullname, address, company_id, phone_number, created_at, updated_at, created_by, updated_by
	FROM customers
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get all customer: %w", err)
	}
	defer rows.Close()

	// Iterate over the rows and scan the results into CustomerModel objects
	customers := make([]*model.CustomerModel, 0)
	for rows.Next() {
		var customer model.CustomerModel
		err := rows.Scan(
			&customer.Id,
			&customer.FullName,
			&customer.Address,
			&customer.CompanyId,
			&customer.PhoneNumber,
			&customer.CreatedAt,
			&customer.UpdatedAt,
			&customer.CreatedBy,
			&customer.UpdatedBy,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan customer row: %w", err)
		}
		customers = append(customers, &customer)
	}

	return customers, nil
}

func (repo *customerRepository) DeleteCustomer(id string) error {
	// Perform database delete operation
	_, err := repo.db.Exec(`
		DELETE FROM customers WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf("failed to delete: %w", err)
	}

	return nil
}
