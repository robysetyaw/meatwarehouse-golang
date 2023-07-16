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

type CashFlowStatement struct {
	StartDate        time.Time
	EndDate          time.Time
	Cash             float64
	TotalPaymentIn   float64
	TotalPaymentOut  float64
	TotalExpenditure float64
	PaymentIn        []*TransactionReportDetail
	PaymentOut       []*TransactionReportDetail
	Expenditure      []*DailyExpenditureReport
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
	PaymentAmount       float64
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

type ConsolidatedReportMaul struct {
	StartDate          time.Time
	EndDate            time.Time
	SalesTotal         float64
	SalesDetail        []*SalesReportOut
	ReceiptTotal       float64
	ReceiptDetail      []*ReceiptReport
	DebtOwnerTotal     float64
	DebtOwnerDetail    []*DebtAccountsPayableReportDetail
	DebtCustomerTotal  float64
	DebtCustomerDetail []*DebtAccountsPayableReportDetail
	ProfitTotal        float64
	ProfitDetail       []*ProfitAndLossStatement
	CashTotal          float64
	CashTotalDetail    []*CashFlowStatement
}
type ConsolidatedReport struct {
	ExpenditureReport         *ExpenditureReport         `json:"expenditure_report"`
	TransactionReport         *TransactionReport         `json:"transaction_report"`
	SalesReport               *SalesReportOut            `json:"sales_report"`
	ReceiptReport             *ReceiptReport             `json:"receipt_report"`
	DebtAccountsPayableReport *DebtAccountsPayableReport `json:"debt_accounts_payable_report"`
	ProfitLossStatement       *ProfitAndLossStatement    `json:"profit_loss_statement"`
	CashFlowStatement         *CashFlowStatement         `json:"cash_flow_statement"`
}

type StockMovementReport struct {
	MeatID        string  `json:"meat_id"`
	MeatName      string  `json:"meat_name"`
	TotalStockIn  float64 `json:"total_stock_in"`
	TotalStockOut float64 `json:"total_stock_out"`
	StockMovement float64 `json:"stock_movement"`
}