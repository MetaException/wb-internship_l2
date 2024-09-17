package main

import (
	"reflect"
	"testing"
)

func TestFindAnagram(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		expected map[string][]string
	}{
		{
			name:     "Basic case with simple anagrams",
			words:    []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			expected: map[string][]string{"пятак": {"пятак", "пятка", "тяпка"}, "листок": {"листок", "слиток", "столик"}},
		},
		{
			name:     "Case with duplicates",
			words:    []string{"пятак", "пятка", "тяпка", "тяпка", "листок", "слиток", "столик"},
			expected: map[string][]string{"пятак": {"пятак", "пятка", "тяпка"}, "листок": {"листок", "слиток", "столик"}},
		},
		{
			name:     "Case with mixed case letters",
			words:    []string{"Пятак", "пятка", "Тяпка", "листок", "Слиток", "Столик"},
			expected: map[string][]string{"пятак": {"пятак", "пятка", "тяпка"}, "листок": {"листок", "слиток", "столик"}},
		},
		{
			name:     "Case with no anagrams",
			words:    []string{"кот", "собака", "дерево"},
			expected: map[string][]string{},
		},
		{
			name:     "Case with empty input",
			words:    []string{},
			expected: map[string][]string{},
		},
		{
			name:     "Case with one word",
			words:    []string{"кот"},
			expected: map[string][]string{},
		},
		{
			name:     "Case with identical words",
			words:    []string{"кот", "кот", "кот"},
			expected: map[string][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindAnagram(tt.words)
			if !reflect.DeepEqual(*result, tt.expected) {
				t.Errorf("expected %v, but got %v", tt.expected, *result)
			}
		})
	}
}
