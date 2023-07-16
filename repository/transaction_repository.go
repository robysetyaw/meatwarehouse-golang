package repository

import (
	"database/sql"
	"fmt"
	"time"

	"enigmacamp.com/final-project/team-4/track-prosto/model"
)

type TransactionRepository interface {
	CreateTransactionHeader(header *model.TransactionHeader) error
	GetTransactionByID(id string) (*model.TransactionHeader, error)
	GetAllTransactions() ([]*model.TransactionHeader, error)
	DeleteTransaction(id string) error
	CountTransactions() (int, error)
	GetByInvoiceNumber(invoice_number string) (*model.TransactionHeader, error)
	UpdateStatusInvoicePaid(id string) error
	UpdateStatusPaymentAmount(id string,total float64 ) error
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (repo *transactionRepository) CreateTransactionHeader(header *model.TransactionHeader) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}
	now := time.Now()
	header.CreatedAt = now
	header.UpdatedAt = now
	header.IsActive = true

	query := "INSERT INTO transaction_headers (id, date, customer_id, name, address, company, phone_number, tx_type, total, is_active, created_at, updated_at, created_by, updated_by, inv_number, payment_amount, payment_status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17) RETURNING id"
	err = tx.QueryRow(query,header.ID, header.Date, header.CustomerID, header.Name, header.Address, header.Company, header.PhoneNumber, header.TxType, header.Total, header.IsActive, header.CreatedAt, header.UpdatedAt, header.CreatedBy, header.UpdatedBy, header.InvoiceNumber, header.PaymentAmount, header.PaymentStatus).Scan(&header.ID)
	if err != nil {
		// tx.Rollback()
		return fmt.Errorf("failed to create transaction header: %w", err)
	}

	// Create transaction details
	query = "INSERT INTO transaction_details (id,transaction_id, meat_id, meat_name, qty, price, total, is_active, created_at, updated_at, created_by, updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)"
	for _, detail := range header.TransactionDetails {
		_, err := tx.Exec(query,detail.ID, header.ID, detail.MeatID, detail.MeatName, detail.Qty, detail.Price, detail.Total, header.IsActive, header.CreatedAt, header.UpdatedAt, header.CreatedBy, header.CreatedBy)
		if err != nil {
			// tx.Rollback()
			return fmt.Errorf("failed to create transaction detail: %w", err)
		}
	}

	updateQuery := "UPDATE transaction_headers SET total = $1 WHERE id = $2"
	_, err = tx.Exec(updateQuery, header.Total, header.ID)
	if err != nil {
		// tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (repo *transactionRepository) GetTransactionByID(id string) (*model.TransactionHeader, error) {
	var transaction model.TransactionHeader

	// Get transaction header from database
	err := repo.db.QueryRow(`
		SELECT id, date, customer_id, name, address, company, phone_number, tx_type, total, is_active, created_at, updated_at, created_by, updated_by, inv_number
		FROM transaction_headers
		WHERE id = $1 AND is_active = true
	`, id).Scan(
		&transaction.ID,
		&transaction.Date,
		&transaction.CustomerID,
		&transaction.Name,
		&transaction.Address,
		&transaction.Company,
		&transaction.PhoneNumber,
		&transaction.TxType,
		&transaction.Total,
		&transaction.IsActive,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
		&transaction.CreatedBy,
		&transaction.UpdatedBy,
		&transaction.InvoiceNumber,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Return nil if transaction header not found
		}
		return nil, fmt.Errorf("failed to get transaction header by ID: %w", err)
	}

	// Get transaction details from database
	rows, err := repo.db.Query(`
		SELECT id, transaction_id, meat_id, meat_name, qty, price, total, is_active, created_at, updated_at, created_by, updated_by
		FROM transaction_details
		WHERE transaction_id = $1 AND is_active = true
	`, transaction.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction details: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var detail model.TransactionDetail
		err := rows.Scan(
			&detail.ID,
			&detail.TransactionID,
			&detail.MeatID,
			&detail.MeatName,
			&detail.Qty,
			&detail.Price,
			&detail.Total,
			&detail.IsActive,
			&detail.CreatedAt,
			&detail.UpdatedAt,
			&detail.CreatedBy,
			&detail.UpdatedBy,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan transaction detail row: %w", err)
		}

		transaction.TransactionDetails = append(transaction.TransactionDetails, &detail)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while iterating over transaction detail rows: %w", err)
	}

	return &transaction, nil
}

func (repo *transactionRepository) GetAllTransactions() ([]*model.TransactionHeader, error) {
	// Perform database query to retrieve all active transactions
	rows, err := repo.db.Query(`
		SELECT id, date, customer_id, name, address, company, phone_number, tx_type, total, is_active, created_at, updated_at, created_by, updated_by, inv_number
		FROM transaction_headers
		WHERE is_active = true
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get all transactions: %w", err)
	}
	defer rows.Close()

	// Iterate over the rows and scan the results into TransactionHeader objects
	transactions := make([]*model.TransactionHeader, 0)
	for rows.Next() {
		var transaction model.TransactionHeader
		err := rows.Scan(
			&transaction.ID,
			&transaction.Date,
			&transaction.CustomerID,
			&transaction.Name,
			&transaction.Address,
			&transaction.Company,
			&transaction.PhoneNumber,
			&transaction.TxType,
			&transaction.Total,
			&transaction.IsActive,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
			&transaction.CreatedBy,
			&transaction.UpdatedBy,
			&transaction.InvoiceNumber,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan transaction header row: %w", err)
		}

		transactions = append(transactions, &transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while iterating over transaction header rows: %w", err)
	}

	return transactions, nil
}

func (repo *transactionRepository) DeleteTransaction(id string) error {
	_, err := repo.db.Exec(`
		UPDATE transaction_headers
		SET is_active = false
		WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf("failed to delete transaction: %w", err)
	}

	return nil
}

func (repo *transactionRepository) CountTransactions() (int, error) {
	var count int

	err := repo.db.QueryRow("SELECT COUNT(*) FROM transaction_headers").Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count transactions: %w", err)
	}

	return count+1, nil
}

func (repo *transactionRepository) GetByInvoiceNumber(invoice_number string) (*model.TransactionHeader, error){
	var transaction model.TransactionHeader

	// Get transaction header from database
	err := repo.db.QueryRow(`
		SELECT id, date, customer_id, name, address, company, phone_number, tx_type, total, is_active, created_at, updated_at, created_by, updated_by, inv_number
		FROM transaction_headers
		WHERE inv_number = $1 AND is_active
	`, invoice_number).Scan(
		&transaction.ID,
		&transaction.Date,
		&transaction.CustomerID,
		&transaction.Name,
		&transaction.Address,
		&transaction.Company,
		&transaction.PhoneNumber,
		&transaction.TxType,
		&transaction.Total,
		&transaction.IsActive,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
		&transaction.CreatedBy,
		&transaction.UpdatedBy,
		&transaction.InvoiceNumber,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Return nil if transaction header not found
		}
		return nil, fmt.Errorf("failed to get transaction header by ID: %w", err)
	}

	// Get transaction details from database
	rows, err := repo.db.Query(`
		SELECT id, transaction_id, meat_id, meat_name, qty, price, total, is_active, created_at, updated_at, created_by, updated_by
		FROM transaction_details
		WHERE transaction_id = $1 AND is_active = true
	`, transaction.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction details: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var detail model.TransactionDetail
		err := rows.Scan(
			&detail.ID,
			&detail.TransactionID,
			&detail.MeatID,
			&detail.MeatName,
			&detail.Qty,
			&detail.Price,
			&detail.Total,
			&detail.IsActive,
			&detail.CreatedAt,
			&detail.UpdatedAt,
			&detail.CreatedBy,
			&detail.UpdatedBy,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan transaction detail row: %w", err)
		}

		transaction.TransactionDetails = append(transaction.TransactionDetails, &detail)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while iterating over transaction detail rows: %w", err)
	}

	return &transaction, nil	
}

func (repo *transactionRepository) UpdateStatusInvoicePaid(id string) error {
	_, err := repo.db.Exec(`
		UPDATE transaction_headers
		SET payment_status = 'paid'
		WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf("failed to update transaction: %w", err)
	}

	return nil
}

func (repo *transactionRepository) UpdateStatusPaymentAmount(id string,total float64 ) error {
	_, err := repo.db.Exec(`
		UPDATE transaction_headers
		SET payment_amount = $1
		WHERE id = $2
	`, total, id)
	if err != nil {
		return fmt.Errorf("failed to update transaction: %w", err)
	}

	return nil
}

