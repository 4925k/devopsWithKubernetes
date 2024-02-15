package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	RANDOMIMAGEURL     = "https://picsum.photos/1200"
	DAILYIMAGEFILEPATH = "/usr/src/app/files/daily-image.jpg"
)

// getDailyImage fetches the daily image to be displayed
// checks whether there is an existing image already
// if yes, displays that image
// otherwise, downloads a daily image from the website.
func getDailyImage() {
	// CHECK IF IMAGE DOES NOT EXIST EXISTS
	if !fileExists(DAILYIMAGEFILEPATH) {
		resp, err := http.Get(RANDOMIMAGEURL)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer resp.Body.Close()

		// CREATE A NEW FILE
		out, err := os.Create(DAILYIMAGEFILEPATH)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer out.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			log.Fatal(err)
			return
		}
	}

}

// fileExists return true is file exists else false
// also returns false if file is older than 24 hours
func fileExists(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}

	// CUSTOME RULE
	// checking if file was created within the last 24 hours
	if f.ModTime().After(time.Now().Add(-24 * time.Hour)) {
		return false
	}

	return true
}
