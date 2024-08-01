package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DBManager struct {
	db *sql.DB
}

func NewManager() (*DBManager, error) {
	connStr := "user=postgres password=mypass dbname=notifydb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	manager := &DBManager{db: db}
	if err := manager.createEventTable(); err != nil {
		return nil, err
	}
	if err := manager.createUserTable(); err != nil {
		return nil, err
	}
	if err := manager.createSubscriptionTable(); err != nil {
		return nil, err
	}
	return manager, nil
}
