package model

import "time"

type ExpenditureReport struct {
	StartDate        time.Time
	EndDate          time.Time
	TotalExpenditure float64
	Expenditures     []*DailyExpenditureReport
}

type DailyExpenditureReport struct {
	ID          string    `json:"id"`
	UserID      string    `json:"-"`
	Username    string    `json:"username"`
	Amount      float64   `json:"amount" binding:"required"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Date        string    `json:"date"`
}

type TransactionReport struct {
	StartDate           time.Time
	EndDate             time.Time
	TotalInTransaction  float64
	TotalOutTransaction float64
	Report              []*TransactionReportDetail
}

type TransactionReportDetail struct {
	InvoiceNumber       string
	CustomerName        string
	CompanyName         string
	PhoneNumberCustomer string
	TxType              string
	Qty                 float64
	Total               float64
}
