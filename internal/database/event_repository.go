package database

func (dm *DBManager) createEventTable() error {

	createTable := `CREATE TABLE IF NOT EXISTS events (
		event_id SERIAL PRIMARY KEY,
		event_name VARCHAR(255) NOT NULL,
		event_date TIMESTAMP NOT NULL
	)`
	_, err := dm.db.Exec(createTable)
	if err != nil {
		return err
	}
	return nil
}
