package core

import (
	"strings"

	"github.com/gnomegl/brutekit/internal/utils"
)

// Transformation map for leet speak substitutions
// This can be customized by adding/modifying character mappings
var Transformations = map[rune][]string{
	'a': {"@", "4"},
	'b': {"8"},
	'c': {"(", "{"},
	'e': {"3"},
	'g': {"9", "6"},
	'h': {"4", "#"},
	'i': {"1", "!", "|"},
	'l': {"1", "|"},
	'o': {"0"},
	's': {"$", "5", "z"},
	't': {"7", "+"},
	'v': {"\\/"},
	'w': {"vv", "uu"},
	'x': {"><"},
	'z': {"2"},
}

// Global mutation storage
var Mutations []string

// GenerateMutations creates all possible variations of a word including:
// - Case variations (upper, lower)
// - Leet speak substitutions
// - Padding variations if enabled
func GenerateMutations(word string) {
	// Basic mutation (just the word)
	Mutations = append(Mutations, word)

	// Case variations
	Mutations = append(Mutations, strings.ToUpper(word))
	Mutations = append(Mutations, strings.ToLower(word))

	// Generate leet speak variations
	leetVariations := GenerateLeetVariations(word)
	Mutations = append(Mutations, leetVariations...)

	// Apply paddings if configured
	if len(utils.CommonPaddings) > 0 {
		var tempMutations []string
		for _, mut := range Mutations {
			for _, pad := range utils.CommonPaddings {
				tempMutations = append(tempMutations, mut+pad)
				tempMutations = append(tempMutations, pad+mut)
			}
		}
		Mutations = append(Mutations, tempMutations...)
	}
}

// GenerateLeetVariations creates all possible leet speak variations
// by recursively substituting characters according to the transformations map
func GenerateLeetVariations(word string) []string {
	var result []string
	result = append(result, word)

	// Try substituting each character
	for i, char := range word {
		if substitutions, exists := Transformations[char]; exists {
			var newVariations []string
			// For each existing variation
			for _, variant := range result {
				// Try each possible substitution
				for _, sub := range substitutions {
					newWord := variant[:i] + sub + variant[i+1:]
					newVariations = append(newVariations, newWord)
				}
			}
			result = append(result, newVariations...)
		}
	}

	return result
}
