package database

import (
	_ "github.com/lib/pq"
)

func (dm *DBManager) createSubscriptionTable() error {
	createTable := `CREATE TABLE IF NOT EXISTS subscriptions (
		id  SERIAL PRIMARY KEY,
		user_id INTEGER,
		event_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(user_id),
		FOREIGN KEY (event_id) REFERENCES events(event_id)
	)`

	_, err := dm.db.Exec(createTable)
	if err != nil {
		return err
	}
	return nil
}
