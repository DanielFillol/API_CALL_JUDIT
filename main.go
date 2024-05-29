package main

import (
	"API_CALL_JUDIT/csv"
	"API_CALL_JUDIT/request"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
	"time"
)

const (
	BATCHSize     = 1000                   // set batch size
	BATCHInterval = 100 * time.Millisecond // set the coll down for each batch
	CollDown      = 500 * time.Millisecond // set the coll down for each single request
	WORKERS       = 10                     // set amount of parallel execution
)

const (
	FILENAME      = "response"          // set single .csv response initial name
	FILEPATH      = "data/requests.csv" // set the path were the requests are stored
	FILESeparator = ','                 // set .csv separator
	SKIPHeader    = true                // set if the .csv reader should skip the file header
)

func main() {
	// Set request variables
	request.WORKERS = WORKERS
	request.BATCHInterval = BATCHInterval
	request.CollDown = CollDown

	// Set log file
	logFile, err := os.Create("output.log.txt")
	if err != nil {
		log.Println("Failed to open log file: " + err.Error())
	}
	defer logFile.Close()

	// Create a multi-writer that writes to both the file and os.Stdout (terminal)
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)

	// Load environment variables from .env file
	err = godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file: ", err)
	}
	// Get the value of the AUTH variable from the environment
	var auth = os.Getenv("AUTH")

	// Load data to be requested from CSV file
	requests, err := csv.Read(FILEPATH, FILESeparator, SKIPHeader)
	if err != nil {
		log.Println("Error loading requests from CSV: ", err)
	}

	// Make API requests asynchronously
	start := time.Now()
	log.Println("Starting API calls...")

	err = request.AllBatchAsync(requests, BATCHSize, auth, FILENAME)
	if err != nil {
		log.Println(err)
	}

	log.Println("All API calls completed in " + time.Since(start).String())
}
