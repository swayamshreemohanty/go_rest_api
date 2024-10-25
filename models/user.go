package models

import (
	"errors"
	"go_rest_api/db"
	"go_rest_api/utils"
)

type User struct {
	ID       	int64   `json:"id"`
	Email   	string 	`binding:"required" json:"email"`
	Password 	string 	`binding:"required" json:"password"`
}

//Save the user to the database
func (u *User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`

	stmt, err := db.DBClient.Prepare(query)
	if err != nil {
		return err
	}

	// Close the statement after the function ends
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	// Execute the query
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	// Get the ID of the inserted row
	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	// Set the ID field of the struct
	u.ID = id

	return nil
}

//Login the user
func (u *User) Login() error {
	query := `SELECT id, email, password FROM users WHERE email = ?`

	stmt, err := db.DBClient.Prepare(query)
	if err != nil {
		return err
	}

	// Close the statement after the function ends
	defer stmt.Close()

	// Execute the query
	row := stmt.QueryRow(u.Email)

	var hashedPassword string

	err = row.Scan(&u.ID, &u.Email, &hashedPassword)

	if err != nil {
		return err
	}

	isPasswordValid := utils.CompareHashAndPassword(hashedPassword,u.Password)

	if !isPasswordValid {
		return errors.New("invalid password")
	}

	return nil
}