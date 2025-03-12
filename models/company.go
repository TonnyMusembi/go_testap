package models

import "time"

type Company struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	LowerName       string    `json:"lower_name"`
	PhysicalAddress string    `json:"physical_address"`
	LoanPeriod      string    `json:"loan_period"`
	ApprovesLoan    bool      `json:"approves_loan"`
	Status          int       `json:"status"`
	Version         string    `json:"version"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
