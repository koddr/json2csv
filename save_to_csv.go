package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

// save provides saving operation for source map.
func (c *config) save(source []map[string]string) error {
	// Create a new file to collect CSV data.
	outputFile, err := os.Create(fmt.Sprintf("%s/%d.csv", c.outputFolder, time.Now().Unix()))
	if err != nil {
		return err
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {

		}
	}(outputFile)

	// Write the header of the CSV file and the successive rows by iterating through
	// the JSON struct array.
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Write the header of the CSV file.
	header := []string{c.contentField, "intent"}
	if err = writer.Write(header); err != nil {
		return err
	}

	// Write collected data to CSV file.
	for _, r := range source {
		// First column – content, second column – intent.
		if err = writer.Write([]string{r["content"], c.qualify(r["content"])}); err != nil {
			return err
		}
	}

	return nil
}
