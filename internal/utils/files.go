package utils

import (
	"bufio"
	_ "embed"
	"os"
	"strings"
)

//go:embed common_padding_values.txt
var EmbeddedPaddings string

// Global storage for common padding values
var CommonPaddings []string

func LoadCommonPaddings() error {
	// Read common paddings from file
	scanner := bufio.NewScanner(strings.NewReader(EmbeddedPaddings))
	for scanner.Scan() {
		val := strings.TrimSpace(scanner.Text())
		if val != "" {
			CommonPaddings = append(CommonPaddings, val)
		}
	}

	return scanner.Err()
}

func WriteResults(filename string, mutations []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, mutation := range mutations {
		if _, err := writer.WriteString(mutation + "\n"); err != nil {
			return err
		}
	}

	return writer.Flush()
}
