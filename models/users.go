package models

import (
	"fmt"
	"main/databasee"
)

type User struct {
	Id        int    `json:"id"`
	Names     string `json:"names"`
	LastNames string `json:"last_names"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Users []User

const UserSchema string = `CREATE TABLE if NOT EXISTS users(
	id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
	names varchar(50) NOT NULL,
	last_names varchar(50) NOT NULL,
	phone varchar(50) NOT NULL,
	email varchar(100) NOT NULL,
	password varchar(100) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

func CreateTableUsers() {
	databasee.ExecuteExec(UserSchema)
}

func (user *User) AddNewUser() {
	fmt.Println(user.Email)
	databasee.ExecuteExec(`INSERT INTO users(names, last_names, phone, email, password) VALUES(?,?,?,?,?)`,
		user.Names, user.LastNames, user.Phone, user.Email, user.Password)
}

func GetUser(id int) *Users {
	rows, err := databasee.ExecuteQuery(`SELECT id,names,last_names,phone,email,password FROM users WHERE id=?`, id)
	if err != nil {
		fmt.Println("Error in GetUser", err)
	}
	defer rows.Close()
	users := &Users{}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Names, &user.LastNames, &user.Phone, &user.Email, &user.Password)
		if err != nil {
			fmt.Println("Error in GetUser to read rows", err)
		}
		*users = append(*users, user)

	}
	return users
}
