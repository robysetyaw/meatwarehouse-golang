package model

import "time"

type CreditPayment struct {
	ID            string    `json:"id"`
	InvoiceNumber string    `json:"inv_number"`
	PaymentDate   string    `json:"payment_date"`
	Amount        float64   `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
	CreatedBy     string    `json:"created_by"`
	UpdatedAt     time.Time `json:"updated_at"`
	UpdatedBy     string    `json:"updated_by"`
}
