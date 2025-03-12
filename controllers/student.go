package controllers

import (
	"log/slog"
	"net/http"
	"strconv"

	"student-api/models"
	"student-api/repositories"

	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	students, err := repositories.GetStudents()
	if err != nil {
		slog.Error("Failed to fetch students", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	slog.Info("Successfully fetched students", "count", len(students))
	c.JSON(http.StatusOK, gin.H{"data": students})
}

func GetStudentByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		slog.Error("Invalid student ID", "id", c.Param("id"), "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	student, err := repositories.GetStudentByID(id)
	if err != nil {
		if err.Error() == "student not found" {
			slog.Warn("Student not found", "id", id)
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			slog.Error("Failed to fetch student", "id", id, "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	slog.Info("Successfully fetched student", "id", id)
	c.JSON(http.StatusOK, gin.H{"data": student})
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		slog.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repositories.CreateStudent(student); err != nil {
		slog.Error("Failed to create student", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	slog.Info("Successfully created student", "student", student)
	c.JSON(http.StatusCreated, gin.H{"data": student})
}

func UpdateStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		slog.Error("Invalid student ID", "id", c.Param("id"), "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		slog.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student.ID = id
	if err := repositories.UpdateStudent(student); err != nil {
		slog.Error("Failed to update student", "id", id, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	slog.Info("Successfully updated student", "id", id)
	c.JSON(http.StatusOK, gin.H{"data": student})
}

func DeleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		slog.Error("Invalid student ID", "id", c.Param("id"), "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	if err := repositories.DeleteStudent(id); err != nil {
		slog.Error("Failed to delete student", "id", id, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	slog.Info("Successfully deleted student", "id", id)
	c.JSON(http.StatusOK, gin.H{"data": "Student deleted successfully"})
}

// func GetStudentsByBranch(c *gin.Context) {
// 	branchID, err := strconv.Atoi(c.Param("branch_id"))
//     if err != nil {
//         slog.Error("Invalid branch ID", "branch_id", c.Param("branch_id"), "error", err)
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid branch ID"})
//         return
//     }

//     students, err := repositories.GetStudentsByBranch(branchID)
//     if err != nil {
//         slog.Error("Failed to fetch students by branch", "branch_id", branchID, "error", err)
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }
//     slog.Info("Successfully fetched students by branch", "branch_id", branchID, "count", len(students))
// }
