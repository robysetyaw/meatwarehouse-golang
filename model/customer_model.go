package model

import "time"

type CustomerModel struct {
	Id          string `json:"id"`
	FullName    string `json:"fullname"`
	Address     string `json:"address"`
	CompanyId   string `json:"company_id"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   time.Time
	UpdateAt    time.Time
	CreatedBy   string
	UpdateBy    string
}
