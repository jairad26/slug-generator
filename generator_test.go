package slug_generator

import (
	"strconv"
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	sg := NewSlugGenerator(1, ADJECTIVES, NOUNS)

	tests := make([]struct {
		name     string
		length   int
		expected int
	}, 100)

	for i := 0; i < 100; i++ {
		tests[i] = struct {
			name     string
			length   int
			expected int
		}{
			name:     "Test " + strconv.Itoa(i+1),
			length:   i + 1,
			expected: i + 1,
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slug := sg.Generate(tt.length)
			if slug == "" {
				t.Errorf("Generate() returned null, expected non-null string")
			}

			words := strings.Split(slug, "-")
			if len(words) != tt.expected {
				t.Errorf("Generate() returned string of length %d, expected %d. Slug: %s", len(words), tt.expected, slug)
			}

			// Rule 1: A slug must end with a noun.
			if !contains(NOUNS, words[len(words)-1]) {
				t.Errorf("Slug does not end with a noun. Slug: %s", slug)
			}

			// Rule 2: You can have at most 2 adjectives describing a noun.
			// Rule 3: A noun is followed by "of" (unless it is the end of the slug).
			// Rule 5: "of" can not be followed by "of".
			adjectiveCount := 0
			for i, word := range words {
				if contains(ADJECTIVES, word) {
					adjectiveCount++
					if adjectiveCount > 2 {
						t.Errorf("More than 2 adjectives(%s, %s) are describing a noun. Slug: %s", words[i-1], words[i], slug)
					}
				} else if contains(NOUNS, word) {
					adjectiveCount = 0
					if i != len(words)-1 && words[i+1] != "of" {
						t.Errorf("Noun (%s) is not followed by 'of'. Slug: %s", word, slug)
					}
				} else if word == "of" {
					if i == len(words)-1 || words[i+1] == "of" {
						t.Errorf("'of' is followed by 'of' or is the last word. Slug: %s", slug)
					}
				}
			}

			// Rule 4: A slug must begin with either a noun or an adjective.
			if !contains(NOUNS, words[0]) && !contains(ADJECTIVES, words[0]) {
				t.Errorf("Slug does not begin with a noun or an adjective. Slug: %s", slug)
			}
		})
	}
}

// Helper function to check if a slice contains a specific string.
func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
