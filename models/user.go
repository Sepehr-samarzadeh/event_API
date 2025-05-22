package models

import (
	"fmt"

	"sep.com/eventapi/db"
	"sep.com/eventapi/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

/*func (u User) Save() error {
	query := "INSERT INTO users(email,password) VALUES(?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.Email, u.Password)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId

	return err
}
*/

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
