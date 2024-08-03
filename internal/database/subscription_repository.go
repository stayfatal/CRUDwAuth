package database

import (
	_ "github.com/lib/pq"
)

type Subscription struct {
	UserID  int `json:"user_id"`
	EventID int `json:"event_id"`
}

func (dm *DBManager) createSubscriptionTable() error {
	createTable := `CREATE TABLE IF NOT EXISTS subscriptions (
		id  SERIAL PRIMARY KEY,
		user_id INTEGER,
		event_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (event_id) REFERENCES events(id)
	)`

	_, err := dm.db.Exec(createTable)
	if err != nil {
		return err
	}
	return nil
}

func (dm *DBManager) AddSubscription(sub Subscription) error {
	_, err := dm.db.Exec("insert into subscriptions (user_id,event_id) values ($1,$2)", sub.UserID, sub.EventID)
	if err != nil {
		return err
	}
	return nil
}
