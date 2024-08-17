package models

import (
	"appointment-tracking/db"
)

type Appointments struct {
	ID         int64  `json:"id"`
	Type       string `json:"type"`
	Client_id  int64  `json:"client_id"`
	Doctors_id int64  `json:"doctors_id"`
}

func (a *Appointments) CreateNewAppointment() error {
	query := `INSERT INTO appointments (type, client_id, doctors_id) VALUES (?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(a.Type, a.Client_id, a.Doctors_id)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	a.ID = id
	return nil
}

func GetAppointmentById(id int64) (*Appointments, error) {
	query := `SELECT * FROM appointments WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var rowFromDb Appointments
	err := row.Scan(&rowFromDb.ID, &rowFromDb.Type, &rowFromDb.Client_id, &rowFromDb.Doctors_id)

	if err != nil {
		return nil, err
	}

	return &rowFromDb, nil
}

func (a *Appointments) UpdateAppointment(id int64) error {

	query := `
	UPDATE appointments SET
	type = ?, client_id = ?, doctors_id = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(a.Type, a.Client_id, a.Doctors_id, id)

	return err

}

func GetAllAppointments() ([]Appointments, error) {
	query := "SELECT * FROM appointments"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var appointments []Appointments

	for rows.Next() {
		var appointemnt Appointments
		err := rows.Scan(&appointemnt.ID, &appointemnt.Type, &appointemnt.Client_id, &appointemnt.Doctors_id)

		if err != nil {
			return nil, err
		}

		appointments = append(appointments, appointemnt)
	}

	return appointments, nil

}

func (a *Appointments) DeleteAppointment() error {
	query := "DELETE FROM appointments WHERE id = ?"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(a.ID)

	return err

}

func GetAppointmentByClientId(id int64) ([]Appointments, error) {
	query := `SELECT id, type, doctors_id FROM appointments WHERE client_id = ?`
	rows, err := db.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appointments []Appointments

	for rows.Next() {
		var appointment Appointments
		err := rows.Scan(&appointment.ID, &appointment.Type, &appointment.Doctors_id)
		if err != nil {
			return nil, err
		}
		appointment.Client_id = id
		appointments = append(appointments, appointment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return appointments, nil
}
