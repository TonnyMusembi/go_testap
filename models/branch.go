package models

import "time"

type Branch struct {
	ID              int       `json:"id"`
	CompanyID       int       `json:"company_id"`
	Name            string    `json:"name"`
	PhysicalAddress string    `json:"physical_address"`
	Status          int       `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Version         string    `json:"version"`
}