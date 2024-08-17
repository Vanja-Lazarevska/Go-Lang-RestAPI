package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {

	db, err := sql.Open("sqlite", "api.sql")
	if err != nil {
		panic("Database could not connect: " + err.Error())
	}

	DB = db

	err = createTables()
	if err != nil {
		panic("Database could not connect: " + err.Error())
	}

	fmt.Println("Tables created successfully!")
}

func createTables() error {
	createStaffTable := `
	CREATE TABLE IF NOT EXISTS staff_members (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
		lastName TEXT NOT NULL UNIQUE,
		role TEXT NOT NULL,
		available BOOLEAN NOT NULL,
		clinic TEXT NOT NULL
	)`
	if _, err := DB.Exec(createStaffTable); err != nil {
		return err
	}

	createClientsTable := `
	CREATE TABLE IF NOT EXISTS clients (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		lastName TEXT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`
	if _, err := DB.Exec(createClientsTable); err != nil {
		return err
	}

	createAppointmentsTable := `
	CREATE TABLE IF NOT EXISTS appointments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		type TEXT NOT NULL,
		client_id INTEGER,
		doctors_id INTEGER,
		FOREIGN KEY(client_id) REFERENCES clients(id),
		FOREIGN KEY(doctors_id) REFERENCES staff_members(id)
	)`
	if _, err := DB.Exec(createAppointmentsTable); err != nil {
		return err
	}

	return nil
}
