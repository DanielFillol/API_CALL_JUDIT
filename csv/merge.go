package csv

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
)

func MergeCSVs(folderPath string, mergedFilename string) error {
	// List all CSV files in the specified folder
	files, err := filepath.Glob(filepath.Join(folderPath, "*.csv"))
	if err != nil {
		return err
	}

	if len(files) == 0 {
		log.Println("No CSV files found in the specified folder. Folder: ", folderPath)
		return nil
	}

	// Create or open the merged CSV file
	mergedFile, err := os.Create(mergedFilename)
	if err != nil {
		return err
	}
	defer mergedFile.Close()

	// Create a CSV writer for the merged file
	mergedWriter := csv.NewWriter(mergedFile)
	defer mergedWriter.Flush()

	// WriteLawsuits headers from the first CSV file
	firstFile, err := os.Open(files[0])
	if err != nil {
		return err
	}
	defer firstFile.Close()

	firstReader := csv.NewReader(firstFile)
	headers, err := firstReader.Read()
	if err != nil {
		return err
	}
	mergedWriter.Write(headers)

	// Iterate through each CSV file, skipping the header in subsequent files
	for _, file := range files {
		if file == mergedFilename {
			continue // Skip the merged file itself
		}

		currentFile, err := os.Open(file)
		if err != nil {
			return err
		}
		defer currentFile.Close()

		currentReader := csv.NewReader(currentFile)
		if _, err := currentReader.Read(); err != nil && err != io.EOF {
			return err
		}

		// WriteLawsuits the remaining rows to the merged file
		lineNumber := 2 // Start from line 2 (skipping header)
		for {
			row, err := currentReader.Read()
			if err == io.EOF {
				break
			}

			// Check if the number of fields in the row matches the number of headers
			if len(row) != len(headers) {
				// If the number of fields is less than headers, pad with empty strings
				if len(row) < len(headers) {
					for i := len(row); i < len(headers); i++ {
						row = append(row, "")
					}
				} else {
					// If the number of fields is more than headers, truncate the record
					row = row[:len(headers)]
				}
			}

			if err := mergedWriter.Write(row); err != nil {
				log.Println(err)
				return err
			}

			lineNumber++
		}
	}

	log.Printf("Merged CSV file created: %s", mergedFilename)

	return nil
}
