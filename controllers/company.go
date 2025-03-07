package controllers

import (
	"log/slog"
	"net/http"
	"student-api/models"
	"student-api/repositories"

	"github.com/gin-gonic/gin"
)

func CreateCompany(c *gin.Context) {
	var company models.Company

	// Bind the JSON request body to the Company struct
	if err := c.ShouldBindJSON(&company); err != nil {
	slog.Error("not json format", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the repository function to create the company
	id, err := repositories.CreateCompany(company)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error("not json format", "error", err)
		return
	}

	// Return the ID of the created company
	c.JSON(http.StatusCreated, gin.H{
		"message": "Company created successfully",
		"id":      id,
	})
}