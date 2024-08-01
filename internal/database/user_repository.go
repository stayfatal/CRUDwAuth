package database

import (
	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Username string
	Password string
}

func (dm *DBManager) createUserTable() error {
	createTable := `CREATE TABLE IF NOT EXISTS users (
		user_id SERIAL PRIMARY KEY,
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
	err := dm.db.QueryRow("insert into users (user_id,username,password) values ($1,$2,$3) RETURNING user_id", user.ID, user.Username, user.Password).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (dm *DBManager) GetUser(id int) (User, error) {
	row := dm.db.QueryRow("select * from users where user_id = $1", id)
	user := User{}
	var ignore string
	err := row.Scan(&user.ID, &user.Username, &ignore)
	if err != nil {
		return User{}, nil
	}
	return user, nil
}
