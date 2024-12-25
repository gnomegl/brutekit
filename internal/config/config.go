package config

import (
	"flag"
	"fmt"
	"os"
)

// Config holds all command line arguments and settings
type Config struct {
	Words                string // Comma separated keywords to mutate
	AppendNumbering      int    // Append numbering range at end of mutations
	NumberingLimit       int    // Max numbering limit (default: 50)
	Years                string // Years to append (single, comma-separated, or range)
	AppendPadding        string // Additional padding values
	CommonPaddingsBefore bool   // Append common paddings before mutations
	CommonPaddingsAfter  bool   // Append common paddings after mutations
	CustomPaddingsOnly   bool   // Use only user provided paddings
	Output               string // Output filename
	Quiet                bool   // Do not print banner
}

func ParseFlags() *Config {
	// Initialize config
	config := &Config{}

	// Define command line flags
	flag.StringVar(&config.Words, "w", "", "Comma separated keywords to mutate")
	flag.IntVar(&config.AppendNumbering, "an", 0, "Append numbering range")
	flag.IntVar(&config.NumberingLimit, "nl", 50, "Numbering limit")
	flag.StringVar(&config.Years, "y", "", "Years to append")
	flag.StringVar(&config.AppendPadding, "ap", "", "Additional padding values")
	flag.BoolVar(&config.CommonPaddingsBefore, "cpb", false, "Common paddings before")
	flag.BoolVar(&config.CommonPaddingsAfter, "cpa", false, "Common paddings after")
	flag.BoolVar(&config.CustomPaddingsOnly, "cpo", false, "Use only custom paddings")
	flag.StringVar(&config.Output, "o", "", "Output filename")
	flag.BoolVar(&config.Quiet, "q", false, "Do not print banner")

	// Parse flags
	flag.Parse()

	// Check required flags
	if config.Words == "" {
		fmt.Printf("\n[%sDebug%s] Words parameter (-w) is required\n", "", "")
		flag.Usage()
		os.Exit(1)
	}

	return config
}
