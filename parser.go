package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	jsoniter "github.com/json-iterator/go"
)

// parse provides the parsing process.
func (c *config) parse() error {
	// Create slice for results.
	results := make([]map[string]string, 0, len(c.jsonFolder.files))

	// Read JSON files and collect data.
	for index, file := range c.jsonFolder.files {
		if file.IsDir() {
			continue
		}

		// Open the given file (without defer, see end of this loop).
		jsonFile, err := os.Open(filepath.Clean(fmt.Sprintf("%s/%s", c.jsonFolder.path, file.Name())))
		if err != nil {
			return err
		}

		// Read data from the given file.
		b, err := io.ReadAll(jsonFile)
		if err != nil {
			return err
		}

		// Create slice for the current result.
		result := make([]map[string]string, 0)

		// Unmarshal JSON data to the result slice.
		err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(b, &result)
		if err != nil {
			return err
		}

		// Filtering and qualifying input data.
		for _, r := range result {
			// Skip unneeded fields and content points and append a current result to the
			// main results slice.
			if c.filter(r[c.contentField]) {
				results = append(results, r)
			}
		}

		// Save data to current CSV chunk.
		if index > 0 && index%c.chunkSize == 0 {
			// Write a new CSV file with collected (and filtered) data.
			if err = c.save(results); err != nil {
				return err
			}

			result = nil  // clean current result slice
			results = nil // clean main results slice
		}

		// Physically close the current open file (without defer) to prevent max open
		// files error.
		err = jsonFile.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
