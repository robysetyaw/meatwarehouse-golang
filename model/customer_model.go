package model

import "time"

type CustomerModel struct {
	Id          string    `json:"id"`
	FullName    string    `json:"fullname"`
	Address     string    `json:"address"`
	CompanyId   string    `json:"company_id"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`

	CompanyDetail Company
}
