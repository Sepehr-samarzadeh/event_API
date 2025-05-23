package models

import (
	"errors"
	"fmt"

	"sep.com/eventapi/db"
	"sep.com/eventapi/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email,password) VALUES(?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("Prepare error:", err)
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		fmt.Println("Exec error:", err)
		return err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		fmt.Println("LastInsertId error:", err)
		return err
	}
	u.ID = userID
	return nil
}

func (u User) ValidateCredentials() error {
	query := "SELECT id , password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword) //link the  output of query to the variable

	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}
	return nil

}
