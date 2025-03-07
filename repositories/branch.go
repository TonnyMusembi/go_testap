package repositories

import (
	"log/slog"
	"student-api/config"
	"student-api/models"
)

func CreateBranch(branch models.Branch) (int64, error) {
	// Log the start of the operation
	slog.Info("Creating branch", "branch", branch)

	// Execute the SQL query
	result, err := config.DB.Exec(`
		INSERT INTO branches (
			company_id, name, physical_address, status, version
		) VALUES (?, ?, ?, ?, ?)`,
		branch.CompanyID, branch.Name, branch.PhysicalAddress, branch.Status, branch.Version,
	)
	if err != nil {
		slog.Error("Failed to create branch", "error", err)
		return 0, err
	}

	// Get the ID of the newly inserted branch
	id, err := result.LastInsertId()
	if err != nil {
		slog.Error("Failed to get last insert ID", "error", err)
		return 0, err
	}

	// Log the successful operation
	slog.Info("Successfully created branch", "id", id)
	return id, nil
}