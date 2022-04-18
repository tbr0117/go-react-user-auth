package model

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

const (
	CheckUserExists         = `SELECT true from users WHERE email = $1`
	LoginQuery              = `SELECT * from users WHERE email = $1`
	UpdateUserPasswordQuery = `UPDATE users SET password = $2 WHERE id = $1`
	DeleteUser              = `DELETE FROM users WHERE email = $1`
	CreateUserQuery         =  `INSERT INTO users(id, name, password, email) VALUES (DEFAULT, $1 , $2, $3);`
	GetUserByIDQuery        =  `SELECT * FROM users WHERE id = $1`
	GetUserByEmailQuery     =  `SELECT * FROM users WHERE email = $1`
)

type User struct {
	ID        string `json:"id,omitempty"`
	Password  string `json:"password,omitempty"`
	Email     string `json:"email,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type UserMethods interface {
	UpdateHashPassword() error
	HasUserExists(dbConn *sql.DB) bool
}


func (user *User) UpdateHashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) HasUserExists(dbConn	*sql.DB) bool {
	rows, err := dbConn.Query(GetUserByEmailQuery, user.Email)
	if err != nil || !rows.Next() {
		return false
	}
	return true
}