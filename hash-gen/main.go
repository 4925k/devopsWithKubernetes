package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	for {

		h, err := bcrypt.GenerateFromPassword([]byte(time.Now().String()), 8)
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Printf("%s %s\n", time.Now().String(), h)

		time.Sleep(time.Second * 5)
	}
}
