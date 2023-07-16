package repository

import (
	"database/sql"
	"fmt"

	"enigmacamp.com/final-project/team-4/track-prosto/model"
)

type CreditPaymentRepository interface {
	CreateCreditPayment(payment *model.CreditPayment) error
	GetAllCreditPayments() ([]*model.CreditPayment, error)
	GetCreditPaymentByID(id string) (*model.CreditPayment, error)
	UpdateCreditPayment(payment *model.CreditPayment) error
	GetTotalCredit(inv_number string) (float64, error)
}

type creditPaymentRepository struct {
	db *sql.DB
}

func NewCreditPaymentRepository(db *sql.DB) CreditPaymentRepository {
	return &creditPaymentRepository{
		db: db,
	}
}

func (repo *creditPaymentRepository) CreateCreditPayment(payment *model.CreditPayment) error {
	// Implementasi create credit payment
	query := "INSERT INTO credit_payments (id, inv_number, payment_date, amount, created_at, created_by, updated_at, updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	_, err := repo.db.Exec(query, payment.ID, payment.InvoiceNumber, payment.PaymentDate, payment.Amount, payment.CreatedAt, payment.CreatedBy, payment.UpdatedAt, payment.UpdatedBy)
	if err != nil {
		return fmt.Errorf("failed to create credit payment: %w", err)
	}
	return nil
}

func (repo *creditPaymentRepository) GetAllCreditPayments() ([]*model.CreditPayment, error) {
	query := "SELECT id, inv_number, payment_date, amount, created_at, created_by, updated_at, updated_by FROM credit_payments"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get credit payments: %w", err)
	}
	defer rows.Close()

	payments := []*model.CreditPayment{}
	for rows.Next() {
		payment := &model.CreditPayment{}
		err := rows.Scan(&payment.ID, &payment.InvoiceNumber, &payment.PaymentDate, &payment.Amount, &payment.CreatedAt, &payment.CreatedBy, &payment.UpdatedAt, &payment.UpdatedBy)
		if err != nil {
			return nil, fmt.Errorf("failed to scan credit payment row: %w", err)
		}
		payments = append(payments, payment)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while iterating over credit payment rows: %w", err)
	}

	return payments, nil
}

func (repo *creditPaymentRepository) GetCreditPaymentByID(id string) (*model.CreditPayment, error) {
	query := "SELECT id, inv_number, payment_date, amount, created_at, created_by, updated_at, updated_by FROM credit_payments WHERE id = $1"
	row := repo.db.QueryRow(query, id)

	payment := &model.CreditPayment{}
	err := row.Scan(&payment.ID, &payment.InvoiceNumber, &payment.PaymentDate, &payment.Amount, &payment.CreatedAt, &payment.CreatedBy, &payment.UpdatedAt, &payment.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("credit payment not found")
		}
		return nil, fmt.Errorf("failed to get credit payment: %w", err)
	}

	return payment, nil
}

func (repo *creditPaymentRepository) UpdateCreditPayment(payment *model.CreditPayment) error {
	query := "UPDATE credit_payments SET inv_number = $1, payment_date = $2, amount = $3, created_at = $4, created_by = $5, updated_at = $6, updated_by = $7 WHERE id = $8"
	_, err := repo.db.Exec(query, payment.InvoiceNumber, payment.PaymentDate, payment.Amount, payment.CreatedAt, payment.CreatedBy, payment.UpdatedAt, payment.UpdatedBy, payment.ID)
	if err != nil {
		return fmt.Errorf("failed to update credit payment: %w", err)
	}

	return nil
}

func (repo *creditPaymentRepository) GetTotalCredit(inv_number string) (float64, error){
	var total float64
	err := repo.db.QueryRow(`
        SELECT SUM(amount) FROM credit_payments
        WHERE inv_number = '$1' 
    `, ).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("failed to get total credit: %w", err)
	}
	return total, nil
}