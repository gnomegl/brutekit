//
// brutekit - Password List Generator
// A powerful mutation-based wordlist generator
//

package main

import (
	"bufio"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"
)

// -------------- Colors -------------- //
// ANSI color codes for terminal output
const (
	MAIN   = "\033[38;5;50m"
	LOGO   = "\033[38;5;41m"
	LOGO2  = "\033[38;5;42m"
	GREEN  = "\033[38;5;82m"
	ORANGE = "\033[0;38;5;214m"
	PRPL   = "\033[0;38;5;26m"
	PRPL2  = "\033[0;38;5;25m"
	RED    = "\033[1;31m"
	END    = "\033[0m"
	BOLD   = "\033[1m"
)

// -------------- Base Settings -------------- //

// Transformation map for leet speak substitutions
// This can be customized by adding/modifying character mappings
var transformations = map[rune][]string{
	'a': {"@", "4"},
	'b': {"8"},
	'e': {"3"},
	'g': {"9", "6"},
	'i': {"1", "!"},
	'o': {"0"},
	's': {"$", "5"},
	't': {"7"},
}

// Config holds all command line arguments and settings
type Config struct {
	words                string // Comma separated keywords to mutate
	appendNumbering      int    // Append numbering range at end of mutations
	numberingLimit       int    // Max numbering limit (default: 50)
	years                string // Years to append (single, comma-separated, or range)
	appendPadding        string // Additional padding values
	commonPaddingsBefore bool   // Append common paddings before mutations
	commonPaddingsAfter  bool   // Append common paddings after mutations
	customPaddingsOnly   bool   // Use only user provided paddings
	output               string // Output filename
	quiet                bool   // Do not print banner
}

// Global mutation storage
var mutations []string
var commonPaddings []string

//go:embed common_padding_values.txt
var embeddedPaddings string

func main() {
	// Parse command line flags and store in config
	config := parseFlags()

	// Print banner if not in quiet mode
	if !config.quiet {
		fmt.Println("\nbrutekit - Password List Generator")
		fmt.Println("A powerful mutation-based wordlist generator")
		fmt.Println("For security testing and password analysis")
	}

	// Load common paddings from file
	if err := loadCommonPaddings(); err != nil {
		fmt.Printf("\n[%sDebug%s] Error loading common paddings: %v\n", RED, END, err)
		os.Exit(1)
	}

	// Add custom paddings to common paddings
	if config.appendPadding != "" {
		for _, val := range strings.Split(config.appendPadding, ",") {
			val = strings.TrimSpace(val)
			if val != "" {
				commonPaddings = append(commonPaddings, val)
			}
		}
	}

	// Process each word
	words := strings.Split(config.words, ",")
	for _, word := range words {
		word = strings.TrimSpace(word)
		if word != "" {
			generateMutations(word)
		}
	}

	// Write results to file
	outputFile := "output.txt"
	if config.output != "" {
		outputFile = config.output
	}

	if err := writeResults(outputFile); err != nil {
		fmt.Printf("\n[%sDebug%s] Error writing results: %v\n", RED, END, err)
		os.Exit(1)
	}

	fmt.Printf("\n[%s+%s] Generated %d mutations\n", GREEN, END, len(mutations))
	fmt.Printf("[%s+%s] Results written to: %s\n", GREEN, END, outputFile)
}

func parseFlags() Config {
	// Initialize config
	config := Config{}

	// Define command line flags
	flag.StringVar(&config.words, "w", "", "Comma separated keywords to mutate")
	flag.IntVar(&config.appendNumbering, "an", 0, "Append numbering range")
	flag.IntVar(&config.numberingLimit, "nl", 50, "Numbering limit")
	flag.StringVar(&config.years, "y", "", "Years to append")
	flag.StringVar(&config.appendPadding, "ap", "", "Additional padding values")
	flag.BoolVar(&config.commonPaddingsBefore, "cpb", false, "Common paddings before")
	flag.BoolVar(&config.commonPaddingsAfter, "cpa", false, "Common paddings after")
	flag.BoolVar(&config.customPaddingsOnly, "cpo", false, "Use only custom paddings")
	flag.StringVar(&config.output, "o", "", "Output filename")
	flag.BoolVar(&config.quiet, "q", false, "Do not print banner")

	// Parse flags
	flag.Parse()

	// Check required flags
	if config.words == "" {
		fmt.Printf("\n[%sDebug%s] Words parameter (-w) is required\n", RED, END)
		flag.Usage()
		os.Exit(1)
	}

	return config
}

func loadCommonPaddings() error {
	// Read common paddings from file
	scanner := bufio.NewScanner(strings.NewReader(embeddedPaddings))
	for scanner.Scan() {
		val := strings.TrimSpace(scanner.Text())
		if val != "" {
			commonPaddings = append(commonPaddings, val)
		}
	}

	return scanner.Err()
}

// -------------- Core Functions -------------- //

// generateMutations creates all possible variations of a word including:
// - Case variations (upper, lower)
// - Leet speak substitutions
// - Padding variations if enabled
func generateMutations(word string) {
	// Basic mutation (just the word)
	mutations = append(mutations, word)

	// Case variations
	mutations = append(mutations, strings.ToUpper(word))
	mutations = append(mutations, strings.ToLower(word))

	// Generate leet speak variations
	leetVariations := generateLeetVariations(word)
	mutations = append(mutations, leetVariations...)

	// Apply paddings if configured
	if len(commonPaddings) > 0 {
		var tempMutations []string
		for _, mut := range mutations {
			for _, pad := range commonPaddings {
				tempMutations = append(tempMutations, mut+pad)
				tempMutations = append(tempMutations, pad+mut)
			}
		}
		mutations = append(mutations, tempMutations...)
	}
}

// generateLeetVariations creates all possible leet speak variations
// by recursively substituting characters according to the transformations map
func generateLeetVariations(word string) []string {
	var results []string
	var chars []rune

	// Convert word to runes for proper UTF-8 handling
	for _, c := range word {
		chars = append(chars, c)
	}

	// Generate all possible combinations
	for i := 0; i < len(chars); i++ {
		if replacements, exists := transformations[chars[i]]; exists {
			for _, repl := range replacements {
				newWord := string(chars[:i]) + repl + string(chars[i+1:])
				results = append(results, newWord)

				// Recursively generate variations for the new word
				subVariations := generateLeetVariations(newWord)
				results = append(results, subVariations...)
			}
		}
	}

	return results
}

// writeResults writes all generated mutations to the specified output file
// Returns an error if writing fails
func writeResults(filename string) error {
	// Create output file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write mutations to file
	writer := bufio.NewWriter(file)
	for _, mutation := range mutations {
		if _, err := writer.WriteString(mutation + "\n"); err != nil {
			return err
		}
	}

	return writer.Flush()
}
