package utils

import (
	"strings"
	"unicode"
)

func RemoveNumbers(line string) string {
	var result strings.Builder

	for _, ch := range line {
		if unicode.IsNumber(ch) {
			result.WriteRune(ch)
		}
	}

	return result.String()
}

func RemoveDuplicates(text []string) []string {
	hashmap := make(map[string]struct{})

	for _, v := range text {
		hashmap[v] = struct{}{}
	}

	result := make([]string, 0)
	for key := range hashmap {
		result = append(result, key)
	}

	return result
}
