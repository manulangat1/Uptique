package status

import (
	"Uptique/internal/common"
	"Uptique/internal/database"
	"Uptique/internal/database/models"
	"Uptique/internal/database/repository"
	"fmt"
	"log"
	"sync"
)

func MonitorStart() error {

	db, err := database.Open()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	websites, err := repository.FetchWebsites(db)

	if err != nil {
		return err
	}

	fmt.Println(websites)

	// now for each website check its status.
	ch := make(chan error, len(websites))
	successCh := make(chan bool, len(websites))
	var wg sync.WaitGroup
	for _, val := range websites {
		wg.Add(1)
		go func(val models.WebsiteModel) {
			defer wg.Done()
			result, err := common.RequestTo(db, &val)

			ch <- err
			successCh <- result

			// ch <- common.RequestTo(db, &val)
		}(val)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	// collect errors
	for err := range ch {
		if err != nil {
			log.Printf("error checking website: %v", err)
		}
	}
	// fetch all status now.

	statuses, err := repository.FetchAllStatus(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("statuses", statuses)

	for _, val := range statuses {
		fmt.Println(val.Status, val.CreatedAt, val.WebsiteID)
	}

	return nil
}
