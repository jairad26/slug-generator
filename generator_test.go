package slug_generator

import (
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	adjectives := []string{"quick", "brown", "lazy"}
	nouns := []string{"fox", "dog", "cat"}
	sg := NewSlugGenerator(1, adjectives, nouns)

	tests := []struct {
		name     string
		length   int
		expected int
	}{
		{"Test 1", 1, 1},
		{"Test 2", 2, 2},
		{"Test 3", 3, 3},
		{"Test 4", 4, 4},
		{"Test 5", 5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slug := sg.Generate(tt.length)
			if slug == "" {
				t.Errorf("Generate() returned null, expected non-null string")
			}

			words := strings.Split(slug, "-")
			if len(words) != tt.expected {
				t.Errorf("Generate() returned string of length %d, expected %d", len(words), tt.expected)
			}
		})
	}
}
