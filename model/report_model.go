package model

import "time"

type ExpenditureReport struct {
	StartDate        time.Time
	EndDate          time.Time
	TotalExpenditure float64
	Expenditures     []*DailyExpenditureReport
}

type DailyExpenditureReport struct {
	No          int
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

type SalesReportOut struct {
	StartDate  time.Time
	EndDate    time.Time
	SalesTotal float64
	Report     []*TransactionReportDetail
}

type ReceiptReport struct {
	StartDate    time.Time
	EndDate      time.Time
	ReceiptTotal float64
	Report       []*TransactionReportDetail
}

type TransactionReportDetail struct {
	No                  int
	InvoiceNumber       string
	Date                string
	CustomerName        string
	CompanyName         string
	PhoneNumberCustomer string
	TxType              string
	Total               float64
	PaymentStatus       string
	DebtTotal           float64
}

type DebtAccountsPayableReport struct {
	StartDate        time.Time
	EndDate          time.Time
	ReceivablesTotal float64
	DebtTotal        float64
	Receivables      []*DebtAccountsPayableReportDetail
	Debt             []*DebtAccountsPayableReportDetail
}

type DebtAccountsPayableReportDetail struct {
	No                  int
	InvoiceNumber       string
	Date                string
	CustomerName        string
	CompanyName         string
	PhoneNumberCustomer string
	TxType              string
	PaymentStatus       string
	Debt                float64
}

type ProfitAndLossStatement struct {
	StartDate          time.Time
	EndDate            time.Time
	SalesTotal         float64
	TotalExpenditure   float64
	TotalInTransaction float64
	Profit             float64
	Sales              []*TransactionReportDetail
	TransactionIn      []*TransactionReportDetail
	Expenditures       []*DailyExpenditureReport
}
