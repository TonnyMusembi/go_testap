package controllers

import (
	"net/http"

	"student-api/models"
	"student-api/repositories"

	"github.com/gin-gonic/gin"
	"strconv"

)

func CreateBranch(c *gin.Context) {
	var branch models.Branch

	// Bind the JSON request body to the Branch struct
	if err := c.ShouldBindJSON(&branch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the repository function to create the branch
	id, err := repositories.CreateBranch(branch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the ID of the created branch
	c.JSON(http.StatusCreated, gin.H{
		"message": "Branch created successfully",
		"id":      id,
	})
}

func GetBranchCount(c *gin.Context) {
	// Parse the company_id from the request
	companyIDStr := c.Param("company_id")
	companyID, err := strconv.Atoi(companyIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}

	// Call the repository function to get the branch count
	count, err := repositories.GetBranchCount(companyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the branch count
	c.JSON(http.StatusOK, gin.H{
		"company_id": companyID,
		"count":      count,
	})
}

func GetBranchesByCompany(c *gin.Context) {
	// Parse the company_id from the request
    companyIDStr := c.Param("company_id")
    companyID, err := strconv.Atoi(companyIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
        return
    }

    // Call the repository function to get the branches by company
    branches, err := repositories.GetBranchesByCompany(companyID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Return the branches by company
    c.JSON(http.StatusOK, gin.H{
        "company_id": companyID,
        "branches":   branches,
    })
}