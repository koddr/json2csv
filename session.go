package main

import (
	"errors"
	"os"
	"path/filepath"
)

// newSession provides init configuration for the parsing session.
func newSession(
	jsonPath, intentsPath, filterPath, outputPath, contentField string,
	minWordLen, chunkSize int,
) (*config, error) {
	if minWordLen < 0 {
		return nil, errors.New("can't parse data with negative word count")
	}

	if chunkSize <= 0 {
		return nil, errors.New("can't create CSV chunk file with zero (or negative) size")
	}

	if contentField == "" {
		return nil, errors.New("can't parse data without content name attribute")
	}

	// Read the given file with intents.
	intentsFile, err := os.ReadFile(filepath.Clean(intentsPath))
	if err != nil {
		return nil, err
	}

	// Read the given file with filter.
	filterFile, err := os.ReadFile(filepath.Clean(filterPath))
	if err != nil {
		return nil, err
	}

	// Read the given folder for output CSV files.
	_, err = os.Stat(filepath.Clean(outputPath))
	if os.IsNotExist(err) {
		// Create a new folder if is not existed.
		if err = os.MkdirAll(filepath.Clean(outputPath), 0750); err != nil {
			return nil, errors.New("can't create folder to output CSV file")
		}
	}

	// Read the given folder with JSON data.
	jsonFiles, err := os.ReadDir(filepath.Clean(jsonPath))
	if err != nil {
		return nil, err
	}

	return &config{
		minWordLen:   minWordLen,
		chunkSize:    chunkSize,
		intentsFile:  intentsFile,
		filterFile:   filterFile,
		outputFolder: filepath.Clean(outputPath),
		contentField: contentField,
		jsonFolder:   jsonFolder{path: jsonPath, files: jsonFiles},
	}, nil
}
