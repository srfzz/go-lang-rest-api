package models

import "go-lang-restapi/db"

type User struct {
	Id       uint64 `json:"id"`
	Email    string `json:"email" bindings:"required,unique"`
	Password string `json:"password" bindings:"required"`
}

func (u *User) Save() error {
	query := "Insert into users(email,password) values(?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}
	userID, err := result.LastInsertId()
	u.Id = uint64(userID)
	return err
}
