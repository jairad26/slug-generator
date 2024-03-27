package slug_generator

import (
	"fmt"
	"math/rand"
	"strings"
)

type SlugGenerator struct {
	random     *rand.Rand
	adjectives []string
	nouns      []string
}

func (sg *SlugGenerator) getRandomAdjective() string {
	return sg.adjectives[rand.Intn(len(sg.adjectives))]
}

func (sg *SlugGenerator) getRandomNoun() string {
	return sg.nouns[rand.Intn(len(sg.nouns))]
}

func NewSlugGenerator(seed int64, adjectives, nouns []string) *SlugGenerator {
	src := rand.NewSource(seed)
	return &SlugGenerator{
		random:     rand.New(src),
		adjectives: adjectives,
		nouns:      nouns,
	}
}

func (sg *SlugGenerator) Generate(length int) string {
	if length < 1 {
		return ""
	}
	if length == 1 {
		return sg.getRandomNoun()
	}
	if length == 2 {
		return fmt.Sprintf("%v-%v", sg.getRandomAdjective(), sg.getRandomNoun())
	}
	if length == 3 {
		return fmt.Sprintf("%v-%v-%v", sg.getRandomAdjective(), sg.getRandomAdjective(), sg.getRandomNoun())
	}

	slugArr := make([]string, length)
	adjectiveCount := 0
	previousWordWasNoun := false
	previousWordWasOf := false

	for i := 0; i < length; i++ {
		if i == length-3 {
			if previousWordWasNoun {
				slugArr[i] = "of"
				previousWordWasOf = true
				previousWordWasNoun = false
				adjectiveCount = 0
			} else if previousWordWasOf {
				// randomly pick noun or adjective
				if sg.random.Intn(2) == 0 {
					slugArr[i] = sg.getRandomNoun()
					previousWordWasNoun = true
					previousWordWasOf = false
					adjectiveCount = 0
				} else {
					slugArr[i] = sg.getRandomAdjective()
					adjectiveCount++
					previousWordWasNoun = false
					previousWordWasOf = false
				}
			} else {
				slugArr[i] = sg.getRandomNoun()
				previousWordWasNoun = true
				previousWordWasOf = false
				adjectiveCount = 0
			}
		} else if i == length-2 {
			if previousWordWasNoun {
				slugArr[i] = "of"
				previousWordWasOf = true
				previousWordWasNoun = false
				adjectiveCount = 0
			} else if previousWordWasOf || adjectiveCount == 1 {
				// 2nd to last word is an adjective, to make room for a noun
				slugArr[i] = sg.getRandomAdjective()
				previousWordWasNoun = false
				previousWordWasOf = false
				adjectiveCount++
			} else {
				// theoretically should never reach here
				// but if it does, slug will be 1 short, which will be caught in tests
				continue
			}
		} else if i == length-1 || adjectiveCount == 2 {
			// Add a noun if two adjectives or last word
			slugArr[i] = sg.getRandomNoun()
			previousWordWasNoun = true
			previousWordWasOf = false
			adjectiveCount = 0
		} else if previousWordWasNoun {
			slugArr[i] = "of"
			previousWordWasOf = true
			previousWordWasNoun = false
			adjectiveCount = 0
		} else {
			if sg.random.Intn(2) == 0 {
				slugArr[i] = sg.getRandomNoun()
				previousWordWasNoun = true
				previousWordWasOf = false
				adjectiveCount = 0
			} else {
				slugArr[i] = sg.getRandomAdjective()
				adjectiveCount++
				previousWordWasNoun = false
				previousWordWasOf = false
			}
		}
	}

	return strings.Join(slugArr, "-")
}
