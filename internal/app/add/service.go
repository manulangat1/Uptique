package add

import (
	// "Uptique/internal/app/common"
	"Uptique/internal/common"
	"Uptique/internal/database"
	"Uptique/internal/database/repository"
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type WebsiteStruct struct {
	URL  string
	Type string
}

func Website(url string) error {
	var websiteType string

	db, err := database.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if url == "" {
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Enter the website URL")

		input, err := reader.ReadString('\n')

		if err != nil {
			return err
		}
		url = strings.TrimSpace(input)

		fmt.Println("Enter the website type")
		inputType, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		websiteType = strings.TrimSpace(inputType)

	}

	fmt.Println("Adding", url)

	var websites []string
	websites = append(websites, url)
	fmt.Println(websiteType)

	website := WebsiteStruct{
		URL:  url,
		Type: websiteType,
	}
	fmt.Println(website)

	newWebsite, err := repository.CreateWebsite(db, url, websiteType)
	if err != nil {
		log.Fatal(err)
	}
	common.RequestTo(db, newWebsite)

	// RequestTo(website.URL)
	// repository.CreateStatusMonitor(db, 1, )

	fmt.Printf("Added website: [%d] %s (%s)\n", newWebsite.ID, newWebsite.URL, newWebsite.Type)

	return nil

}

func RequestTo(url string) {
	d := time.Now()
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	// Always close the response body to close leaks.
	defer res.Body.Close()
	// Read and print the body.
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	duration := time.Since(d)
	fmt.Println(duration / time.Second)
	fmt.Println(string(body))
}
