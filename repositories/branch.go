package repositories

import (
	"log/slog"
	"strconv"
	"student-api/config"
	"student-api/models"
	"golang.org/x/sync/singleflight"
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

var (
	sfGroup singleflight.Group // SingleFlight group to manage concurrent requests
)

func GetBranchCount(companyID int) (int, error) {
	// Use SingleFlight to deduplicate requests
	key := "x" + strconv.Itoa(companyID) // Unique key for the query
	result, err, _ := sfGroup.Do(key, func() (interface{}, error) {
		// Perform the database query
		var count int
		err := config.DB.QueryRow("SELECT COUNT(*) FROM branches WHERE company_id = ?", companyID).Scan(&count)
		if err != nil {
			slog.Error("Failed to query branch count", "company_id", companyID, "error", err)
			return 0, err
		}
		slog.Info("Successfully queried branch count", "company_id", companyID, "count", count)
		return count, nil
	})

	if err != nil {
		return 0, err
	}

	return result.(int), nil
}	