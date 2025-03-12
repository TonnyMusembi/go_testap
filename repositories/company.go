package repositories

import (
	"log/slog"
	"student-api/config"
	"student-api/models"
)

func CreateCompany(company models.Company) (int64, error) {
	// Log the start of the operation
	slog.Info("Creating company", "company", company)

	// Execute the SQL query
	result, err := config.DB.Exec(`
		INSERT INTO companies (
			name, lower_name, physical_address, loan_period, approves_loan, status, version
		) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		company.Name, company.LowerName, company.PhysicalAddress, company.LoanPeriod, company.ApprovesLoan, company.Status, company.Version,
	)
	if err != nil {
		slog.Error("Failed to create company", "error", err)
		return 0, err
	}

	// Get the ID of the newly inserted company
	id, err := result.LastInsertId()
	if err != nil {
		slog.Error("Failed to get last insert ID", "error", err)
		return 0, err
	}

	// Log the successful operation
	slog.Info("Successfully created company", "id", id)
	return id, nil
}

func GetCompanyByID(id int) (models.Company, error) {
	// Log the start of the operation
    slog.Info("Fetching company by ID", "id", id)

    var company models.Company
    err := config.DB.QueryRow(
	"SELECT id, name, lower_name, physical_address, loan_period, approves_loan, status, version FROM companies WHERE id =?", id).Scan(&company.ID, &company.Name, &company.LowerName, &company.PhysicalAddress, &company.LoanPeriod, &company.ApprovesLoan, &company.Status, &company.Version)
    if err != nil {
        slog.Error("Failed to fetch company", "id", id, "error", err)
        return company, err
    }

    // Log the successful operation
    slog.Info("Successfully fetched company", "id", id)
    return company, nil

}
