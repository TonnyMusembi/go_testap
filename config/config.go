package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDatabase() {
    dsn := "root:tonny@07@tcp(127.0.0.1:3306)/users"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        panic(fmt.Sprintf("Failed to connect to database: %s", err.Error()))
    }

    if err = db.Ping(); err != nil {
        panic(fmt.Sprintf("Failed to ping database: %s", err.Error()))
    }

    DB = db

    // Create table if it doesn't exist
    query := `CREATE TABLE IF NOT EXISTS students (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        age INT NOT NULL,
        grade VARCHAR(50) NOT NULL
    );`
    if _, err := DB.Exec(query); err != nil {
        panic(fmt.Sprintf("Failed to create table: %s", err.Error()))
    }
}