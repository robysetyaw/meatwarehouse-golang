package repository

import (
	"database/sql"
	"fmt"
	"time"

	"enigmacamp.com/final-project/team-4/track-prosto/model"
)

type CompanyRepository interface {
	CreateCompany(*model.Company) error
	UpdateCompany(*model.Company) error
	GetCompanyById(string) (*model.Company, error)
	GetCompanyByName(string) (*model.Company, error)
	GetAllCompany() ([]*model.Company, error)
	DeleteCompany(string) error
}

type companyRepository struct {
	db *sql.DB
}

func NewCompanyRepository(db *sql.DB) CompanyRepository {
	return &companyRepository{
		db: db,
	}
}

func (repo *companyRepository) CreateCompany(company *model.Company) error {
	now := time.Now()
	company.CreatedAt = now
	company.UpdatedAt = now
	// Perform database insert operation
	_, err := repo.db.Exec(`
		INSERT INTO companies (id, company_name, address, email, phone_number, is_active, created_at, updated_at, created_by, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`, company.ID, company.CompanyName, company.Address, company.Email, company.PhoneNumber, company.IsActive, company.CreatedAt, company.UpdatedAt, company.CreatedBy, company.CreatedBy)
	if err != nil {
		return fmt.Errorf("failed to create company: %w", err)
	}

	return nil
}

func (repo *companyRepository) UpdateCompany(company *model.Company) error {
	// Set updated_at timestamp
	company.UpdatedAt = time.Now()

	// Perform database update operation
	_, err := repo.db.Exec(`
		UPDATE companies
		SET company_name = $1, address = $2, email = $3, phone_number = $4, is_active = $5, updated_at = $6, updated_by = $7
		WHERE id = $8
	`, company.CompanyName, company.Address, company.Email, company.PhoneNumber, company.IsActive, company.UpdatedAt, company.UpdatedBy, company.ID)
	if err != nil {
		return fmt.Errorf("failed to update customer: %w", err)
	}

	return nil
}

func (repo *companyRepository) GetCompanyById(id string) (*model.Company, error) {
	var company model.Company

	// Perform database query to retrieve the company by ID
	err := repo.db.QueryRow(`
		SELECT id, company_name, address, email, phone_number, is_active, created_at, updated_at, created_by, updated_by
		FROM companies
		WHERE id = $1
	`, id).Scan(
		&company.ID,
		&company.CompanyName,
		&company.Address,
		&company.Email,
		&company.PhoneNumber,
		&company.IsActive,
		&company.CreatedAt,
		&company.UpdatedAt,
		&company.CreatedBy,
		&company.UpdatedBy,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // company not found
		}
		return nil, fmt.Errorf("failed to get company by ID: %w", err)
	}

	return &company, nil
}
func (repo *companyRepository) GetCompanyByName(companyName string) (*model.Company, error) {
	var company model.Company

	// Perform database query to retrieve the company by ID
	err := repo.db.QueryRow(`
		SELECT id, company_name, address, email, phone_number, is_active, created_at, updated_at, created_by, updated_by
		FROM companies
		WHERE company_name = $1
	`, companyName).Scan(
		&company.ID,
		&company.CompanyName,
		&company.Address,
		&company.Email,
		&company.PhoneNumber,
		&company.IsActive,
		&company.CreatedAt,
		&company.UpdatedAt,
		&company.CreatedBy,
		&company.UpdatedBy,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // company not found
		}
		return nil, fmt.Errorf("failed to get company by company_name: %w", err)
	}

	return &company, nil
}

func (repo *companyRepository) GetAllCompany() ([]*model.Company, error) {
	// Perform database query to retrieve all companies
	rows, err := repo.db.Query(`
	SELECT id, company_name, address, email, phone_number, is_active, created_at, updated_at, created_by, updated_by
	FROM companies
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get all company: %w", err)
	}
	defer rows.Close()

	// Iterate over the rows and scan the results into Company objects
	companies := make([]*model.Company, 0)
	for rows.Next() {
		var company model.Company
		err := rows.Scan(
			&company.ID,
			&company.CompanyName,
			&company.Address,
			&company.Email,
			&company.PhoneNumber,
			&company.IsActive,
			&company.CreatedAt,
			&company.UpdatedAt,
			&company.CreatedBy,
			&company.UpdatedBy,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan company row: %w", err)
		}
		companies = append(companies, &company)
	}

	return companies, nil
}

func (repo *companyRepository) DeleteCompany(id string) error {
	// Perform database delete operation
	_, err := repo.db.Exec(`
		DELETE FROM companies WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf("failed to delete: %w", err)
	}

	return nil
}
