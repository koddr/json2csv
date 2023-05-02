package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	// Catch flags.
	jsonPath := flag.String("json", "./json_files", "set path to folder with JSON files or single file with input data")
	intentsPath := flag.String("intents", "./intents.json", "set path to JSON file with intents dictionary")
	filterPath := flag.String("filter", "./filter.json", "set path to JSON file with filter dictionary")
	outputPath := flag.String("output", "./output_files", "set path to folder for output CSV files")
	contentField := flag.String("content-field", "content", "set name of the content field in your JSON struct")
	minWordLen := flag.Int("min-word-len", 0, "set min length of word to parse from JSON files")
	chunkSize := flag.Int("chunk", 5000, "set chunk size for output CSV file")

	// Parse flags.
	flag.Parse()

	// Print the welcome message.
	fmt.Println("Hello and welcome to json2csv ✌️")
	fmt.Println("\nConfig for this session is creating now. Please wait...")

	// Create a new parse session with the given configs from flags.
	session, err := newSession(*jsonPath, *intentsPath, *filterPath, *outputPath, *contentField, *minWordLen, *chunkSize)
	if err != nil {
		log.Fatal(err)
	}

	// Print the configuration message.
	fmt.Println("\nOK! 👌 This is your config:")
	fmt.Printf("\n – folder with input JSON files is '%s'\n", *jsonPath)
	fmt.Printf(" – file with intents dictionary is '%s'\n", *intentsPath)
	fmt.Printf(" – file with filter dictionary is '%s'\n", *filterPath)
	fmt.Printf(" – folder for output CSV files is '%s'\n", *outputPath)
	fmt.Printf(" – name of the content field is '%s'\n", *contentField)
	fmt.Printf(" – min length of word to parse is %d letters\n", *minWordLen)
	fmt.Printf(" – chunk size for output CSV file is %d (per file)\n", *chunkSize)
	fmt.Println("\nParser is starting now. Please wait...")

	// Start timer.
	start := time.Now()

	// Run parser with a given config.
	if err := session.parse(); err != nil {
		log.Fatal(err)
	}

	// Stop timer and print results.
	fmt.Printf("All done! ✨ Time elapsed: %v\n", time.Since(start))
}
