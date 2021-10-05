package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID			string	`json:"id,omitempty"`
	Email		string	`json:"email,omitempty"`
	Password	string	`json:"password,omitempty"`
	Confirm		string	`json:"confirm,omitempty"`
	FirstName	string	`json:"firstname,omitempty"`
	LastName	string	`json:"lastname,omitempty"`
	CreatedAt	int64	`json:"created_at,omitempty"`
	UpdatedAt	int64	`json:"updated_at,omitempty"`
}
const (
	TABLE = "users"
)
type UserModel struct {
	DB		*sql.DB
}

func (model *UserModel) GetUserByEmail(email string) (bool, User) {
	var user User
	row := model.DB.QueryRow(fmt.Sprintf("SELECT id, email, password FROM %s WHERE email = $1",TABLE), email)
	if row.Err() != nil {
		return false, user
	}
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return false, user
	}
	return true, user
}

func (model *UserModel) SaveUser(user *User) (string, error) {
	stmt, err := model.DB.Prepare(fmt.Sprintf("INSERT INTO %s(id, firstname, lastname, email, password, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7)", TABLE))
	if err != nil {
		return "", err
	}
	_, err = stmt.Exec(user.ID, user.FirstName, user.LastName, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}

func (model *UserModel) GetUserById(id string) (User, error){
	var user User
	row := model.DB.QueryRow(fmt.Sprintf("SELECT id, email, firstname, lastname, created_at FROM %s WHERE id = $1",TABLE), id)
	if row.Err() != nil {
		return user, row.Err()
	}
	err := row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.CreatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}