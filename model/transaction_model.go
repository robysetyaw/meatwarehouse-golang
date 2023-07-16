package model

import "time"

type TransactionHeader struct {
	ID                 string               `json:"id"`
	Date               string               `json:"date"`
	InvoiceNumber      string               `json:"invoice_number"`
	CustomerID         string               `json:"customer_id"`
	Name               string               `json:"name"`
	Address            string               `json:"address"`
	Company            string               `json:"company"`
	PhoneNumber        string               `json:"phone_number"`
	TxType             string               `json:"tx_type"`
	PaymentStatus      string               `json:"payment_status"`
	PaymentAmount      float64              `json:"payment_amount"`
	Total              float64              `json:"total"`
	IsActive           bool                 `json:"is_active"`
	CreatedAt          time.Time            `json:"created_at"`
	UpdatedAt          time.Time            `json:"updated_at"`
	CreatedBy          string               `json:"created_by"`
	UpdatedBy          string               `json:"updated_by"`
	TransactionDetails []*TransactionDetail `json:"transaction_details"`
}

type TransactionDetail struct {
	ID            string    `json:"id"`
	TransactionID string    `json:"transaction_id"`
	MeatID        string    `json:"meat_id"`
	MeatName      string    `json:"meat_name"`
	Qty           float64   `json:"qty"`
	Price         float64   `json:"price"`
	Total         float64   `json:"total"`
	IsActive      bool      `json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedBy     string    `json:"created_by"`
	UpdatedBy     string    `json:"updated_by"`
}

func (h *TransactionHeader) CalulatedTotal() {
	total := 0.0
	for _, detail := range h.TransactionDetails {
		detail.Total = detail.Price * detail.Qty
		total += detail.Total
	}
	h.Total = total
}
