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

	for {
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Println("no db connection", err)
			time.Sleep(20 * time.Second)
			continue
		}

		err := db.Ping()
		if err != nil {
			log.Println("no db connection while pinging", err)
			time.Sleep(20 * time.Second)
			continue
		}

		break
	}

	dbSetup()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", 3001),
		Handler: routes(),
	}

	log.Print("starting server")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("server started at %v\n", srv.Addr)
}
