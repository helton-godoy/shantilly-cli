package tui

import (
	"strings"
)

// removeString removes a string from a slice
func removeString(slice []string, s string) []string {
	for i, item := range slice {
		if item == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// ParseOption exports parseOptions for external use
func ParseOption(options []string) []string {
	return parseOptions(options)
}

// parseOptions converts comma-separated options to array
func parseOptions(options []string) []string {
	var result []string
	for _, opt := range options {
		parts := strings.Split(opt, ",")
		for _, part := range parts {
			part = strings.TrimSpace(part)
			if part != "" {
				result = append(result, part)
			}
		}
	}
	return result
}
