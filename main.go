//
// brutekit - Password List Generator
// A powerful mutation-based wordlist generator
//

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gnomegl/brutekit/internal/config"
	"github.com/gnomegl/brutekit/internal/core"
	"github.com/gnomegl/brutekit/internal/utils"
)

func main() {
	// Parse command line flags and store in config
	cfg := config.ParseFlags()

	// Print banner if not in quiet mode
	if !cfg.Quiet {
		fmt.Println("\nbrutekit - Password List Generator")
		fmt.Println("A powerful mutation-based wordlist generator")
		fmt.Println("For security testing and password analysis")
	}

	// Load common paddings from file
	if err := utils.LoadCommonPaddings(); err != nil {
		fmt.Printf("\n[%sDebug%s] Error loading common paddings: %v\n", utils.RED, utils.END, err)
		os.Exit(1)
	}

	// Add custom paddings to common paddings
	if cfg.AppendPadding != "" {
		for _, val := range strings.Split(cfg.AppendPadding, ",") {
			val = strings.TrimSpace(val)
			if val != "" {
				utils.CommonPaddings = append(utils.CommonPaddings, val)
			}
		}
	}

	// Process each word
	words := strings.Split(cfg.Words, ",")
	for _, word := range words {
		word = strings.TrimSpace(word)
		if word != "" {
			core.GenerateMutations(word)
		}
	}

	// Write results to file
	outputFile := "output.txt"
	if cfg.Output != "" {
		outputFile = cfg.Output
	}

	if err := utils.WriteResults(outputFile, core.Mutations); err != nil {
		fmt.Printf("\n[%sDebug%s] Error writing results: %v\n", utils.RED, utils.END, err)
		os.Exit(1)
	}

	fmt.Printf("\n[%s+%s] Generated %d mutations\n", utils.GREEN, utils.END, len(core.Mutations))
	fmt.Printf("[%s+%s] Results written to: %s\n", utils.GREEN, utils.END, outputFile)
}
