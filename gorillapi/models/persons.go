package models

// Person struct for person
type Person struct {
	ID    int64  `json:"id"`
	Login string `json:"login"`
}

// PersonList gets persons from database
func PersonList() ([]Person, error) {
	sqliteDB, err := Connect()
	defer sqliteDB.Close()

	if err != nil {
		return nil, err
	}

	rows, err := sqliteDB.Query(`SELECT id, login FROM person`)

	if err != nil {
		return nil, err
	}

	var persons []Person

	for rows.Next() {
		var person Person

		rows.Scan(&person.ID, &person.Login)
		persons = append(persons, person)
	}

	return persons, nil
}
