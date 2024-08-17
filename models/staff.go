package models

import (
	"appointment-tracking/db"
)

type Staff struct {
	ID        int64
	Name      string `binding:"required"`
	LastName  string `binding:"required"`
	Role      string `binding:"required"`
	Available bool   `binding:"required"`
	Clinic    string `binding:"required"`
}

func (s *Staff) CreateStaff() error {
	query := `INSERT INTO staff_members(name, lastName,role, available, clinic) 
	VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(s.Name, s.LastName, s.Role, s.Available, s.Clinic)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	s.ID = id
	return err
}

func GetAllStaff() ([]Staff, error) {
	query := "SELECT * FROM staff_members"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var staffs []Staff

	for rows.Next() {
		var staff Staff
		err := rows.Scan(&staff.ID, &staff.Name, &staff.LastName, &staff.Role, &staff.Available, &staff.Clinic)

		if err != nil {
			return nil, err
		}

		staffs = append(staffs, staff)
	}

	return staffs, nil

}

func GetStaffById(id int64) (*Staff, error) {
	query := `SELECT * FROM staff_members WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var rowFromDb Staff
	err := row.Scan(&rowFromDb.ID, &rowFromDb.Name, &rowFromDb.LastName, &rowFromDb.Role, &rowFromDb.Available, &rowFromDb.Clinic)

	if err != nil {
		return nil, err
	}

	return &rowFromDb, nil
}

func (s *Staff) UpdateDoctor() error {

	query := `
	UPDATE staff_members SET
	name = ?, lastName = ?, role = ?, available = ?, clinic = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(s.Name, s.LastName, s.Role, s.Available, s.Clinic, s.ID)

	return err

}

func (s *Staff) DeleteDoctor() error {
	query := "DELETE FROM staff_members WHERE id = ?"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(s.ID)

	return err
}
