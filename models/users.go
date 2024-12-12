package models

import (
	"errors"

	"gitub.com/Jidetireni/events-restapi/db"
	"gitub.com/Jidetireni/events-restapi/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES(?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userid, err := result.LastInsertId()
	u.ID = userid
	return err
}

func (u User) ValidateCred() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return err
	}

	passwordIsvalid := utils.ComparePassword(u.Password, retrievedPassword)

	if !passwordIsvalid {
		return errors.New("invalid credentials")
	}
	return nil
}

// func (u User) GetUserId() error {
// 	query := "SELECT id FROM users WHERE email = ?"

// 	row := db.DB.QueryRow(query, u.Email)
// 	err := row.Scan(&u.ID)

// 	return err

// }
