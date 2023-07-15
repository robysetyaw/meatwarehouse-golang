package repository

import (
	"database/sql"
	"fmt"

	"enigmacamp.com/final-project/team-4/track-prosto/model"
)

type TransactionRepository interface {
	CreateTransaction(transaction *model.TransactionHeader) error
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (repo *transactionRepository) CreateTransaction(transaction *model.TransactionHeader) error {
	// Insert transaction header
	queryHeader := `
		INSERT INTO transaction_headers (date, customer_id, name, address, company, phone_number, tx_type, total, is_active, created_at, updated_at, created_by, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING id
	`

	err := repo.db.QueryRow(queryHeader,
		transaction.Date,
		transaction.CustomerID,
		transaction.Name,
		transaction.Address,
		transaction.Company,
		transaction.PhoneNumber,
		transaction.TxType,
		transaction.Total,
		transaction.IsActive,
		transaction.CreatedAt,
		transaction.UpdatedAt,
		transaction.CreatedBy,
		transaction.UpdatedBy,
	).Scan(&transaction.ID)

	if err != nil {
		return fmt.Errorf("failed to create transaction header: %w", err)
	}

	// Insert transaction details
	queryDetail := `
		INSERT INTO transaction_details (transaction_id, meat_id, meat_name, qty, price, total, is_active, created_at, updated_at, created_by, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	for _, detail := range transaction.TransactionDetails {
		_, err := repo.db.Exec(queryDetail,
			transaction.ID,
			detail.MeatID,
			detail.MeatName,
			detail.Qty,
			detail.Price,
			detail.Total,
			detail.IsActive,
			detail.CreatedAt,
			detail.UpdatedAt,
			detail.CreatedBy,
			detail.UpdatedBy,
		)

		if err != nil {
			return fmt.Errorf("failed to create transaction detail: %w", err)
		}
	}

	return nil
}
