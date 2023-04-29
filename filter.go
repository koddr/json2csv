package main

import (
	"strings"
	"unicode/utf8"

	jsoniter "github.com/json-iterator/go"
)

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

	for key, words := range skipWords {
		if key == "skip_prefix" && len(words) > 0 {
			for _, word := range words {
				if strings.HasPrefix(strings.ToLower(s), strings.ToLower(word)) {
					return false
				}
			}
		}
	}

	for key, words := range skipWords {
		if key == "skip_suffix" && len(words) > 0 {
			for _, word := range words {
				if strings.HasSuffix(strings.ToLower(s), strings.ToLower(word)) {
					return false
				}
			}
		}
	}

	for key, words := range skipWords {
		if key == "skip_words" && len(words) > 0 {
			for _, word := range words {
				if strings.Contains(strings.ToLower(s), strings.ToLower(word)) {
					return false
				}
			}
		}
	}

	return true
}
