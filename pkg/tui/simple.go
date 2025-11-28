package tui

import (
	"fmt"
	"strings"
)

// Simple implementations for immediate functionality

// ShowDialog simple version
func ShowDialog(title string, choices []string, defaultVal string, multi bool) ([]string, error) {
	if len(choices) == 0 {
		return nil, fmt.Errorf("no choices provided")
	}

	fmt.Printf("\n%s\n", title)
	for i, choice := range choices {
		prefix := " "
		if choice == defaultVal {
			prefix = "*"
		}
		fmt.Printf("%s %d. %s\n", prefix, i+1, choice)
	}

	fmt.Print("\nSelect option (number): ")
	var input string
	fmt.Scanln(&input)

	// Simple validation
	for i, choice := range choices {
		if fmt.Sprintf("%d", i+1) == input {
			return []string{choice}, nil
		}
	}

	return []string{choices[0]}, nil
}

// ShowForm simple version
func ShowForm(title string, fieldDefs []string, defaultValues []string) (map[string]string, error) {
	if len(fieldDefs) == 0 {
		return nil, fmt.Errorf("no fields provided")
	}

	result := make(map[string]string)

	fmt.Printf("\n%s\n", title)
	for i, fieldDef := range fieldDefs {
		parts := strings.Split(fieldDef, ":")
		if len(parts) < 2 {
			continue
		}

		fieldName := parts[0]

		prompt := fieldName
		if strings.Contains(fieldDef, "required") {
			prompt += " *"
		}

		defaultVal := ""
		if i < len(defaultValues) {
			defaultVal = defaultValues[i]
		}

		if defaultVal != "" {
			fmt.Printf("%s [%s]: ", prompt, defaultVal)
		} else {
			fmt.Printf("%s: ", prompt)
		}

		var input string
		fmt.Scanln(&input)

		if input == "" && defaultVal != "" {
			input = defaultVal
		}

		result[fieldName] = input
	}

	return result, nil
}

// ShowProgress simple version
func ShowProgress(steps []string, current int) error {
	if len(steps) == 0 {
		return fmt.Errorf("no steps provided")
	}

	if current < 0 {
		current = 0
	} else if current >= len(steps) {
		current = len(steps) - 1
	}

	fmt.Printf("\nProgress: %.1f%%\n", float64(current+1)/float64(len(steps))*100)

	for i, step := range steps {
		prefix := "○"
		if i < current {
			prefix = "✓"
		} else if i == current {
			prefix = "◉"
		}
		fmt.Printf("%s %s\n", prefix, step)
	}

	fmt.Printf("\nCurrent: %s\n", steps[current])
	return nil
}

// ShowFileSelector simple version
func ShowFileSelector(path, filter string, multi bool) ([]string, error) {
	fmt.Printf("\nFile Selector: %s\n", path)
	fmt.Printf("Filter: %s\n", filter)
	fmt.Print("Enter file path: ")

	var input string
	fmt.Scanln(&input)

	if input == "" {
		return []string{}, fmt.Errorf("no file selected")
	}

	return []string{input}, nil
}

// ShowConfirm simple version
func ShowConfirm(message, defaultVal string) (string, error) {
	if message == "" {
		return "", fmt.Errorf("message cannot be empty")
	}

	fmt.Printf("\n%s\n", message)

	defaultDisplay := "Y/n"
	if defaultVal == "no" || defaultVal == "n" {
		defaultDisplay = "y/N"
	}

	fmt.Printf("[%s] ", defaultDisplay)
	var input string
	fmt.Scanln(&input)

	input = strings.ToLower(input)
	if input == "" {
		input = defaultVal
	}

	if input == "y" || input == "yes" {
		return "yes", nil
	} else {
		return "no", nil
	}
}
