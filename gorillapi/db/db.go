package db

import (
	"database/sql"
)

// Connect provide connection to local sqlite database.
func Connect() (*sql.DB, error) {
	return sql.Open("sqlite3", "./foo.db")
}

// Init create fake database
func Init() error {
	sqliteDB, err := Connect()
	defer sqliteDB.Close()

	if err != nil {
		return err
	}

	stmt, _ := sqliteDB.Prepare(`CREATE TABLE IF NOT EXISTS person (id INTEGER PRIMARY KEY, login TEXT)`)
	_, err = stmt.Exec()

	if err != nil {
		return err
	}

	stmt, _ = sqliteDB.Prepare(`DELETE FROM person`)
	_, err = stmt.Exec()

	if err != nil {
		return err
	}

	stmt, _ = sqliteDB.Prepare(`CREATE TABLE IF NOT EXISTS article (id INTEGER PRIMARY KEY, person INTEGER, article TEXT)`)
	_, err = stmt.Exec()

	if err != nil {
		return err
	}

	stmt, _ = sqliteDB.Prepare(`DELETE FROM article`)
	_, err = stmt.Exec()

	if err != nil {
		return err
	}

	stmt, _ = sqliteDB.Prepare(`INSERT INTO person(id, login) VALUES(?,?)`)
	_, err = stmt.Exec(1, "jamie")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(2, "eldrich")

	if err != nil {
		return err
	}

	stmt, _ = sqliteDB.Prepare(`INSERT INTO article (id, person, article) VALUES(?,?,?)`)
	_, err = stmt.Exec(1, 1, "Jamie article.")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(2, 2, "Eldrich article")

	if err != nil {
		return err
	}

	return nil
}
