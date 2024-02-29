package main

import (
	"fmt"
	"log"
)

func setPingpongTable() {
	stmt := `CREATE TABLE IF NOT EXISTS pingpong (
		endpoint text,
		count int
	)`

	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}

	stmt = `INSERT INTO pingpong (endpoint, count) VALUES ("pingpong", 0); `

	_, err = db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}

}

func IncreaseCounter() {
	stmt := `UPDATE pingpong  SET count=count+1 WHERE endpoint='pingpong'`

	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}
}

func GetCount() int {
	stmt := `SELECT count FROM pingpong WHERE endpoint='pingpong'`

	var count int

	db.QueryRow(stmt).Scan(&count)

	fmt.Println(count)

	return count
}
