package repositories

import (
	"database/sql"
	"errors"
	"log/slog"
	"student-api/config"
	"student-api/models"
)

func GetStudents() ([]models.Student, error) {
	// Log the start of the operation
	slog.Info("Fetching all students")

	rows, err := config.DB.Query("SELECT id, name, age, grade FROM students")
	if err != nil {
		slog.Error("Failed to fetch students", "error", err)
		return nil, err
	}
	defer rows.Close()

	students := []models.Student{}
	for rows.Next() {
		var student models.Student
		if err := rows.Scan(&student.ID, &student.Name, &student.Age, &student.Grade); err != nil {
			slog.Error("Failed to scan student row", "error", err)
			return nil, err
		}
		students = append(students, student)
	}

	// Log the successful operation
	slog.Info("Successfully fetched students", "count", len(students))
	return students, nil
}

func GetStudentByID(id int) (models.Student, error) {
	// Log the start of the operation
	slog.Info("Fetching student by ID", "id", id)

	var student models.Student
	err := config.DB.QueryRow("SELECT id, name, age, grade FROM students WHERE id = ?", id).Scan(&student.ID, &student.Name, &student.Age, &student.Grade)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Warn("Student not found", "id", id)
			return student, errors.New("student not found")
		}
		slog.Error("Failed to fetch student by ID", "id", id, "error", err)
		return student, err
	}

	// Log the successful operation
	slog.Info("Successfully fetched student", "id", id)
	return student, nil
}

func CreateStudent(student models.Student) error {
	// Log the start of the operation
	slog.Info("Creating student", "student", student)

	_, err := config.DB.Exec("INSERT INTO students (name, age, grade) VALUES (?, ?, ?)", student.Name, student.Age, student.Grade)
	if err != nil {
		slog.Error("Failed to create student", "student", student, "error", err)
		return err
	}

	// Log the successful operation
	slog.Info("Successfully created student", "student", student)
	return nil
}

func UpdateStudent(student models.Student) error {
	// Log the start of the operation
	slog.Info("Updating student", "id", student.ID)

	_, err := config.DB.Exec("UPDATE students SET name = ?, age = ?, grade = ? WHERE id = ?", student.Name, student.Age, student.Grade, student.ID)
	if err != nil {
		slog.Error("Failed to update student", "id", student.ID, "error", err)
		return err
	}

	// Log the successful operation
	slog.Info("Successfully updated student", "id", student.ID)
	return nil
}

func DeleteStudent(id int) error {
	// Log the start of the operation
	slog.Info("Deleting student", "id", id)

	_, err := config.DB.Exec("DELETE FROM students WHERE id = ?", id)
	if err != nil {
		slog.Error("Failed to delete student", "id", id, "error", err)
		return err
	}

	// Log the successful operation
	slog.Info("Successfully deleted student", "id", id)
	return nil
}