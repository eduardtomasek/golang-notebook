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

// Get return person by ID
func (p *Person) Get(id int64) error {
	sqliteDB, err := Connect()
	defer sqliteDB.Close()

	if err != nil {
		return err
	}

	rows, err := sqliteDB.Query(`SELECT login FROM person WHERE id = $1`, id)

	if err != nil {
		return err
	}

	for rows.Next() {
		rows.Scan(&p.Login)
	}

	p.ID = id

	return nil
}

// Add creates new user
func (p *Person) Add(id int64, login string) error {
	sqliteDB, err := Connect()
	defer sqliteDB.Close()

	if err != nil {
		return err
	}

	return nil
}
