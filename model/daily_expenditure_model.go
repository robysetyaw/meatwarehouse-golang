package model

import "time"

type DailyExpenditure struct {
	ID          string    `json:"id"`
	User_ID     string    `json:"user_id"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
}
