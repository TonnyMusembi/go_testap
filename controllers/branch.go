package controllers

import (
	"net/http"

	"student-api/models"
	"student-api/repositories"

	"github.com/gin-gonic/gin"
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