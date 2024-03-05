package main

import "log"

var (
	createTable = `CREATE TABLE IF NOT EXISTS todolist (
		id SERIAL PRIMARY KEY,
		description text
	)`
)

func dbSetup() {
	_, err := db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}

func GetAllTodo() ([]string, error) {
	stmt := `SELECT description FROM todolist`

	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []string
	for rows.Next() {
		var todo string
		err := rows.Scan(&todo)
		if err != nil {
			return nil, err
		}

		out = append(out, todo)
	}

	return out, err
}
