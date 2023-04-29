package main

import (
	"strings"
	"unicode/utf8"

	jsoniter "github.com/json-iterator/go"
)

func (c *config) qualify(s string) string {
	// Skip empty strings.
	if s == "" {
		return ""
	}

	// Skip strings less than min word length.
	if utf8.RuneCountInString(s) < c.minWordLen {
		return ""
	}

	// Create a new map for JSON.
	qualifyWords := make(map[string][]string, 0)

	// Unmarshal JSON data to the result slice.
	err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(c.intentsFile, &qualifyWords)
	if err != nil {
		return ""
	}

	for intent, words := range qualifyWords {
		for _, el := range words {
			if strings.Contains(strings.ToLower(s), strings.ToLower(el)) {
				return intent
			}
		}
	}

	return ""
}
