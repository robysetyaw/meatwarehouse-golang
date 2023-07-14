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
