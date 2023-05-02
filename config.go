package main

import "os"

type config struct {
	minWordLen, chunkSize      int
	intentsFile, filterFile    []byte
	outputFolder, contentField string
	jsonFolder                 jsonFolder
}

type jsonFolder struct {
	path  string
	files []os.DirEntry
}
