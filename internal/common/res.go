package common

import (
	"Uptique/internal/database/models"
	"Uptique/internal/database/repository"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func RequestTo(db *sql.DB, website *models.WebsiteModel) {
	var status string
	d := time.Now()
	res, err := http.Get(website.URL)
	if err != nil {
		log.Fatal(err)
		return
	}
	// Always close the response body to close leaks.
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		status = "up"
	} else {
		status = "down"
	}

	// Read and print the body.
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	duration := time.Since(d)
	fmt.Println(duration / time.Second)
	fmt.Println(string(body))
	new, err := repository.CreateStatusMonitor(db, int(website.ID), status)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Added the status monitor for ", new.WebsiteID)
}
