package database

import (
	"time"
)

type Event struct {
	ID   int       `json:"id"`
	Name string    `json:"name"`
	Date time.Time `json:"timestamp"`
}

type Broadcast struct {
	EventName string
	UserEmail string
}

func (dm *DBManager) createEventTable() error {

	createTable := `CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		date TIMESTAMP NOT NULL
	)`
	_, err := dm.db.Exec(createTable)
	if err != nil {
		return err
	}
	return nil
}

func (dm *DBManager) AddEvent(event Event) error {
	_, err := dm.db.Exec("insert into events (name,date) values ($1,$2)", event.Name, event.Date.Format("2006-01-02T15:04:05Z"))
	if err != nil {
		return err
	}
	return nil
}

func (dm *DBManager) DeleteEvent() {

}

func (dm *DBManager) GetOccuredEvents() ([]Broadcast, error) {
	query := `
        SELECT events.name, users.email
        FROM events
        JOIN subscriptions ON events.id = subscriptions.id
        JOIN users ON subscriptions.user_id = users.id
        WHERE events.date > $1`

	rows, err := dm.db.Query(query, time.Now().Format("2006-01-02T15:04:05Z"))

	var rowArr []Broadcast
	for rows.Next() {
		var temp Broadcast
		rows.Scan(&temp.EventName, &temp.UserEmail)
		rowArr = append(rowArr, temp)
	}
	if err != nil {
		return nil, err
	}
	return rowArr, nil
}
