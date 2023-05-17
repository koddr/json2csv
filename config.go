package main

import "os"

// config presents struct for the configuration.
type config struct {
	outputFolder string
	contentField string
	jsonFolder   jsonFolder
	intentsFile  []byte
	filterFile   []byte
	minWordLen   int
	chunkSize    int
}

// jsonFolder presents struct for the folder with JSON sources.
type jsonFolder struct {
	path  string
	files []os.DirEntry
}
