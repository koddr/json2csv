package main

import "os"

type config struct {
	outputFolder string
	contentField string
	jsonFolder   jsonFolder
	intentsFile  []byte
	filterFile   []byte
	minWordLen   int
	chunkSize    int
}

type jsonFolder struct {
	path  string
	files []os.DirEntry
}
