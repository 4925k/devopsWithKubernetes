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
