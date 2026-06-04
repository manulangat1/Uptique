package status

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
)

func PingWebsite(url string, db *sql.DB) (bool, error) {

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return false, nil
	}
	// Always close the response body to close leaks.
	defer res.Body.Close()
	// Read and print the body.
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return false, nil
	}

	fmt.Println(body)
	return true, nil

}
