package common

import (
	"Uptique/internal/database/models"
	"Uptique/internal/database/repository"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
)

func RequestTo(db *sql.DB, website *models.WebsiteModel) (bool, error) {
	var status string
	res, err := http.Get(website.URL)
	if err != nil {
		log.Fatal(err)
		return false, err
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
		return false, err
	}
	// duration := time.Since(d)
	// fmt.Println(duration / time.Second)

	new, err := repository.CreateStatusMonitor(db, int(website.ID), status)
	if err != nil {
		log.Fatal(err, body)
		return false, err
	}
	fmt.Println("Added the status monitor for ", new.WebsiteID)
	return true, nil
}
