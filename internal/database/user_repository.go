package database

import (
	_ "github.com/lib/pq"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (dm *DBManager) createUserTable() error {
	createTable := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) NOT NULL,
		username VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL
	)`
	_, err := dm.db.Exec(createTable)
	if err != nil {
		return err
	}
	return nil
}

func (dm *DBManager) CreateUser(user User) (int, error) {
	var id int
	err := dm.db.QueryRow("insert into users (email,username,password) values ($1,$2,$3) RETURNING id", user.Email, user.Username, user.Password).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (dm *DBManager) GetUser(id int) (User, error) {
	row := dm.db.QueryRow("select * from users where id = $1", id)
	user := User{}
	var ignore string
	err := row.Scan(&user.ID, &user.Email, &user.Username, &ignore)
	if err != nil {
		return User{}, nil
	}
	return user, nil
}
