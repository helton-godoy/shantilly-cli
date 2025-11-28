package tui

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Enhanced implementations with better parsing

// EnhancedShowDialog with better option parsing
func EnhancedShowDialog(title string, options []string, defaultVal string, multi bool) ([]string, error) {
	parsedOptions := parseOptions(options)
	if len(parsedOptions) == 0 {
		return nil, fmt.Errorf("no valid options provided")
	}

	fmt.Printf("\n=== %s ===\n", title)
	for i, choice := range parsedOptions {
		prefix := " "
		if choice == defaultVal {
			prefix = "*"
		}
		fmt.Printf("%s %d. %s\n", prefix, i+1, choice)
	}

	if multi {
		fmt.Printf("\nSelect options (numbers separated by spaces): ")
		var input string
		fmt.Scanln(&input)

		var selected []string
		parts := strings.Fields(input)
		for _, part := range parts {
			if num, err := strconv.Atoi(part); err == nil && num >= 1 && num <= len(parsedOptions) {
				selected = append(selected, parsedOptions[num-1])
			}
		}
		return selected, nil
	} else {
		fmt.Printf("\nSelect option (number): ")
		var input string
		fmt.Scanln(&input)

		if num, err := strconv.Atoi(input); err == nil && num >= 1 && num <= len(parsedOptions) {
			return []string{parsedOptions[num-1]}, nil
		}

		// Fallback to default or first option
		if defaultVal != "" {
			return []string{defaultVal}, nil
		}
		return []string{parsedOptions[0]}, nil
	}
}

// EnhancedShowForm with validation
func EnhancedShowForm(title string, fieldDefs []string, defaultValues []string) (map[string]string, error) {
	if len(fieldDefs) == 0 {
		return nil, fmt.Errorf("no fields provided")
	}

	result := make(map[string]string)

	fmt.Printf("\n=== %s ===\n", title)

	for i, fieldDef := range fieldDefs {
		parts := strings.Split(fieldDef, ":")
		if len(parts) < 2 {
			continue
		}

		fieldName := parts[0]
		fieldType := parts[1]

		prompt := fieldName
		if strings.Contains(fieldDef, "required") {
			prompt += " *"
		}

		defaultVal := ""
		if i < len(defaultValues) {
			defaultVal = defaultValues[i]
		}

		// Show field type hint
		typeHint := ""
		switch fieldType {
		case "email":
			typeHint = " (email)"
		case "password":
			typeHint = " (password)"
		case "number":
			typeHint = " (number)"
		}

		for {
			if defaultVal != "" {
				fmt.Printf("%s%s [%s]: ", prompt, typeHint, defaultVal)
			} else {
				fmt.Printf("%s%s: ", prompt, typeHint)
			}

			var input string
			fmt.Scanln(&input)

			if input == "" && defaultVal != "" {
				input = defaultVal
			}

			// Validation
			if strings.Contains(fieldDef, "required") && input == "" {
				fmt.Printf("  ⚠️  This field is required. Please enter a value.\n")
				continue
			}

			// Type validation
			switch fieldType {
			case "email":
				if !strings.Contains(input, "@") || !strings.Contains(input, ".") {
					fmt.Printf("  ⚠️  Please enter a valid email address.\n")
					continue
				}
			case "number":
				if _, err := strconv.Atoi(input); err != nil {
					fmt.Printf("  ⚠️  Please enter a valid number.\n")
					continue
				}
			}

			result[fieldName] = input
			break
		}
	}

	return result, nil
}

// EnhancedShowProgress with animation
func EnhancedShowProgress(steps []string, current int) error {
	if len(steps) == 0 {
		return fmt.Errorf("no steps provided")
	}

	if current < 0 {
		current = 0
	} else if current >= len(steps) {
		current = len(steps) - 1
	}

	percent := float64(current+1) / float64(len(steps)) * 100
	barWidth := 30
	filled := int(float64(barWidth) * percent / 100)
	empty := barWidth - filled

	fmt.Printf("\n=== Progress ===\n")
	fmt.Printf("Progress: %.1f%%\n", percent)
	fmt.Printf("[%s%s]\n", strings.Repeat("█", filled), strings.Repeat("░", empty))

	for i, step := range steps {
		prefix := "○"
		if i < current {
			prefix = "✅"
		} else if i == current {
			prefix = "🔄"
		}
		fmt.Printf("%s %s\n", prefix, step)
	}

	if current < len(steps) {
		fmt.Printf("\n📍 Current: %s\n", steps[current])
		if current < len(steps)-1 {
			fmt.Printf("⏭️  Next: %s\n", steps[current+1])
		}
	} else {
		fmt.Printf("\n🎉 All steps completed!\n")
	}

	return nil
}

// EnhancedShowFileSelector with better navigation
func EnhancedShowFileSelector(path, filter string, multi bool) ([]string, error) {
	fmt.Printf("\n=== File Selector ===\n")
	fmt.Printf("📁 Path: %s\n", path)
	if filter != "" {
		fmt.Printf("🔍 Filter: %s\n", filter)
	}

	fmt.Printf("\nEnter file path (or use '.' for current directory): ")
	var input string
	fmt.Scanln(&input)

	if input == "" {
		return []string{}, fmt.Errorf("no file selected")
	}

	// Simple path validation
	if strings.TrimSpace(input) == "." {
		wd, err := os.Getwd()
		if err != nil {
			return []string{input}, nil
		}
		return []string{wd}, nil
	}

	return []string{input}, nil
}

// EnhancedShowConfirm with better UX
func EnhancedShowConfirm(message, defaultVal string) (string, error) {
	if message == "" {
		return "", fmt.Errorf("message cannot be empty")
	}

	fmt.Printf("\n=== Confirmation ===\n")
	fmt.Printf("❓ %s\n", message)

	defaultDisplay := "Y/n"
	defaultIcon := "✅"
	if defaultVal == "no" || defaultVal == "n" {
		defaultDisplay = "y/N"
		defaultIcon = "❌"
	}

	fmt.Printf("\n[%s] %s: ", defaultDisplay, defaultIcon)
	var input string
	fmt.Scanln(&input)

	input = strings.ToLower(strings.TrimSpace(input))
	if input == "" {
		input = defaultVal
	}

	switch input {
	case "y", "yes", "sim", "s":
		fmt.Printf("✅ Confirmed: YES\n")
		return "yes", nil
	case "n", "no", "não", "nao":
		fmt.Printf("❌ Confirmed: NO\n")
		return "no", nil
	default:
		fmt.Printf("❓ Using default: %s\n", strings.ToUpper(defaultVal))
		return defaultVal, nil
	}
}
