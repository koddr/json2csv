package main

import (
	"strings"
	"unicode/utf8"

	jsoniter "github.com/json-iterator/go"
)

// filter provides rule sets for the input string to skip specified values.
func (c *config) filter(s string) bool {
	// Skip empty strings.
	if s == "" {
		return false
	}

	// Skip strings less than min word length.
	if utf8.RuneCountInString(s) < c.minWordLen {
		return false
	}

	// Create a new map for JSON.
	skipWords := make(map[string][]string, 0)

	// Unmarshal JSON data to the result slice.
	err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(c.filterFile, &skipWords)
	if err != nil {
		return false
	}

	// Filter the given words.
	for key, words := range skipWords {
		switch key {
		case "skip_prefixes":
			for _, word := range words {
				if strings.HasPrefix(strings.ToLower(s), strings.ToLower(word)) {
					return false
				}
			}
		case "skip_suffixes":
			for _, word := range words {
				if strings.HasSuffix(strings.ToLower(s), strings.ToLower(word)) {
					return false
				}
			}
		case "skip_words":
			for _, word := range words {
				if strings.Contains(strings.ToLower(s), strings.ToLower(word)) {
					return false
				}
			}
		}
	}

	return true
}
