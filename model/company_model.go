package model

import "time"

type Company struct {
	ID          string    `json:"id"`
	CompanyName string    `json:"company_name" binding:"required"`
	Address     string    `json:"address"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
}
