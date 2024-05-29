package request

import (
	"API_CALL_JUDIT/csv"
	"API_CALL_JUDIT/models"
	"log"
	"strconv"
	"time"
)

const (
	FolderMerge    = "data/responses/"
	MergedFilename = "merged_result.csv"
)

var WORKERS int
var BATCHInterval time.Duration
var CollDown time.Duration

func AllBatchAsync(requests []models.ReadCsv, batchSize int, auth string, fileName string) error {
	err := batchCall(requests, batchSize, auth, FolderMerge, FolderMerge, fileName, MergedFilename)
	if err != nil {
		log.Print("Error on criminal caller: ", err)
	}

	return nil
}

func batchCall(requests []models.ReadCsv, batchSize int, auth string, folderMerge string, folderName string, fileName string, mergedFileName string) error {
	var urlCaller = "https://requests.prod.judit.io/requests"

	// Process requests in batches
	var resultsSaved int
	for i := 0; i < len(requests); i += batchSize {
		end := i + batchSize
		if end > len(requests) {
			end = len(requests)
		}

		batchRequests := requests[i:end]

		// Make API requests asynchronously
		start := time.Now()
		log.Printf("Starting API calls batch %d...", i/batchSize)

		batchResults, err := AsyncAPIRequest(batchRequests, WORKERS, urlCaller, auth, CollDown)
		if err != nil {
			log.Printf("Error making API requests for batch %d: %v", i/batchSize, err)
			continue
		}

		log.Printf("Finished API calls batch %d in %v", i/batchSize, time.Since(start))

		// WriteLawsuits API response to CSV file
		err = csv.WriteLawsuits(fileName+"_"+strconv.Itoa(i/batchSize), folderName, batchResults)
		if err != nil {
			log.Printf("Error writing API response to CSV for batch %d: %v", i/batchSize, err)
			continue
		}

		resultsSaved += len(batchResults)
		// Introduce interval between batches
		time.Sleep(BATCHInterval)
	}

	if resultsSaved != 0 {
		err := csv.MergeCSVs(folderMerge, mergedFileName)
		if err != nil {
			return err
		}
	}

	return nil
}
