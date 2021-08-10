package repository

import (
	"api/src/model"
	"database/sql"
)

func InsertUser(db *sql.DB, user model.User) (*sql.Result, error) {
	stmt, erro := db.Prepare("insert into users (name, nick, email, password) values (?,?,?,?)")
	if erro != nil {
		return nil, erro
	}
	defer stmt.Close()

	result, erro := stmt.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil {
		return nil, erro
	}

	return &result, nil
}

func GetUsers(db *sql.DB) ([]model.User, error) {
	result, erro := db.Query("select * from users")
	if erro != nil {
		return nil, erro
	}
	defer result.Close()

	var users []model.User

	for result.Next() {
		var user model.User
		if erro = result.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.Password, &user.CreatedAt); erro != nil {
			return nil, erro
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUserById(db *sql.DB, ID int) (model.User, error) {
	result, erro := db.Query("select * from users u where u.id = ?", ID)
	if erro != nil {
		return model.User{}, erro
	}
	defer result.Close()

	var user model.User
	for result.Next() {
		if erro = result.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.Password, &user.CreatedAt); erro != nil {
			return model.User{}, erro
		}
	}

	return user, nil
}
