package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var currentStatus string

func main() {
	// generate a hash and time every 5 second
	go func() {
		for {
			// GENERATING A NEW HASH
			h, err := bcrypt.GenerateFromPassword([]byte(time.Now().String()), 8)
			if err != nil {
				log.Fatal(err)
				return
			}

			// OUTPUT FORMATTING
			// '<time> <hash>'
			currentStatus = fmt.Sprintf("%s %s\n", time.Now().String(), h)

			// OUTPUT TO STDOUT
			fmt.Println(currentStatus)

			// OUTPUT TO FILE.

			// COMMENTED part will append to existing file
			// open file
			// file, err := os.OpenFile("hash.txt", os.O_CREATE|os.O_WRONLY, 0644)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// defer file.Close()

			// // write to file
			// _, err = file.WriteString(currentStatus)
			// if err != nil {
			// 	log.Fatal(err)
			// }

			// this is overwrite any existing file
			os.WriteFile("/usr/src/app/files/hash.txt", []byte(currentStatus), fs.FileMode(os.O_CREATE))

			// SLEEP
			// to generate hash every 5 second
			time.Sleep(time.Second * 5)
		}
	}()

	select {}
}
