package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func main() {
	connStr := "user=postgres dbname=postgres password=mysecretpassword sslmode=disable"

	log.Print("here")
	// database
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	// setPingpongTable()

	log.Print("table setup completed")

	// server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", "4445"),
		Handler: routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("pingpong started on %s", srv.Addr)
}
