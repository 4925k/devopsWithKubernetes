package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func main() {
	connStr := "host=postgres user=postgres dbname=postgres password=admin sslmode=disable"

	log.Print("here")
	// database
	for {
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Println("no connection", err)
			time.Sleep(20 * time.Second)
			continue
		}

		break
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
