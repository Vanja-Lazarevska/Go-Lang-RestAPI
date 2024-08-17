package models

import (
	"appointment-tracking/db"
	"appointment-tracking/utils"
	"errors"
)

type Client struct {
	ID           int64          `json:"id"`
	Name         string         `json:"name"`
	LastName     string         `json:"lastName"`
	Email        string         `json:"email" binding:"required"`
	Password     string         `json:"password" binding:"required"`
	Appointments []Appointments `json:"appointments,omitempty"`
}

func (c *Client) CreateNewClient() error {
	query := "INSERT INTO clients (name, lastName, email, password) VALUES (?, ?, ?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(c.Password)

	if err != nil {
		return err
	}
	result, err := stmt.Exec(c.Name, c.LastName, c.Email, hashedPassword)

	if err != nil {
		return err
	}

	clientId, err := result.LastInsertId()

	if err != nil {
		return err
	}

	c.ID = clientId

	return nil
}

func (c *Client) ValidateCredentials() error {
	query := "SELECT id, password FROM clients WHERE email = ? "
	row := db.DB.QueryRow(query, c.Email)

	var retrivedPassword string
	err := row.Scan(&c.ID, &retrivedPassword)

	if err != nil {
		return errors.New("Credentials are invalid")
	}

	passwordIsValid := utils.ComparePasswords(c.Password, retrivedPassword)

	if !passwordIsValid {
		return errors.New("Credentials are invalid")
	}

	return nil

}

func GetAllClients() ([]Client, error) {
	query := "SELECT id, name, lastName, email, password FROM clients"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []Client

	for rows.Next() {
		var client Client
		err := rows.Scan(&client.ID, &client.Name, &client.LastName, &client.Email, &client.Password)
		if err != nil {
			return nil, err
		}

		appointments, err := GetAppointmentByClientId(client.ID)
		if err != nil {
			return nil, err
		}

		client.Appointments = appointments
		clients = append(clients, client)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return clients, nil
}
