package main

import (
	"fmt"
	"os"

	"github.com/helton-godoy/shantilly-cli/pkg/tui"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "shantilly",
		Short: "Modern TUI for shell scripts",
		Long:  "Shantilly-CLI provides beautiful terminal interfaces for shell scripts",
	}

	// Add subcommands
	rootCmd.AddCommand(createDialogCmd())
	rootCmd.AddCommand(createFormCmd())
	rootCmd.AddCommand(createProgressCmd())
	rootCmd.AddCommand(createSelectCmd())
	rootCmd.AddCommand(createConfirmCmd())
	rootCmd.AddCommand(createBubbleTeaCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func createDialogCmd() *cobra.Command {
	var title string
	var options []string
	var defaultVal string
	var multi bool

	var cmd = &cobra.Command{
		Use:   "dialog",
		Short: "Create interactive dialog boxes",
		Long:  "Create beautiful interactive dialog boxes for user selection",
		Run: func(cmd *cobra.Command, args []string) {
			if title == "" || len(options) == 0 {
				fmt.Fprintf(os.Stderr, "Error: --title and --options are required\n")
				os.Exit(1)
			}

			result, err := tui.EnhancedShowDialog(title, options, defaultVal, multi)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			for _, selection := range result {
				fmt.Println(selection)
			}
		},
	}

	cmd.Flags().StringVar(&title, "title", "", "Dialog title (required)")
	cmd.Flags().StringArrayVar(&options, "options", []string{}, "Comma-separated options (required)")
	cmd.Flags().StringVar(&defaultVal, "default", "", "Default selection")
	cmd.Flags().BoolVar(&multi, "multi", false, "Allow multiple selections")

	return cmd
}

func createFormCmd() *cobra.Command {
	var title string
	var fields []string
	var values []string

	var cmd = &cobra.Command{
		Use:   "form",
		Short: "Create data collection forms",
		Long:  "Create forms for collecting user input with validation",
		Run: func(cmd *cobra.Command, args []string) {
			if len(fields) == 0 {
				fmt.Fprintf(os.Stderr, "Error: --fields is required\n")
				os.Exit(1)
			}

			result, err := tui.EnhancedShowForm(title, fields, values)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			for field, value := range result {
				fmt.Printf("%s=%s\n", field, value)
			}
		},
	}

	cmd.Flags().StringVar(&title, "title", "", "Form title")
	cmd.Flags().StringArrayVar(&fields, "fields", []string{}, "Field definitions (required)")
	cmd.Flags().StringArrayVar(&values, "values", []string{}, "Default values")

	return cmd
}

func createProgressCmd() *cobra.Command {
	var steps []string
	var current int

	var cmd = &cobra.Command{
		Use:   "progress",
		Short: "Show progress indicators",
		Long:  "Display beautiful progress bars for long-running operations",
		Run: func(cmd *cobra.Command, args []string) {
			if len(steps) == 0 {
				fmt.Fprintf(os.Stderr, "Error: --steps is required\n")
				os.Exit(1)
			}

			err := tui.EnhancedShowProgress(steps, current)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringArrayVar(&steps, "steps", []string{}, "Progress steps (required)")
	cmd.Flags().IntVar(&current, "current", 0, "Current step index")

	return cmd
}

func createSelectCmd() *cobra.Command {
	var path string
	var filter string
	var multi bool

	var cmd = &cobra.Command{
		Use:   "select",
		Short: "File and directory selection",
		Long:  "Interactive file and directory browser with filtering",
		Run: func(cmd *cobra.Command, args []string) {
			if path == "" {
				path = "."
			}

			result, err := tui.EnhancedShowFileSelector(path, filter, multi)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			for _, file := range result {
				fmt.Println(file)
			}
		},
	}

	cmd.Flags().StringVar(&path, "path", "", "Starting path (default: current directory)")
	cmd.Flags().StringVar(&filter, "filter", "", "File filter pattern")
	cmd.Flags().BoolVar(&multi, "multi", false, "Allow multiple selections")

	return cmd
}

func createConfirmCmd() *cobra.Command {
	var message string
	var defaultVal string

	var cmd = &cobra.Command{
		Use:   "confirm",
		Short: "Yes/No confirmation dialogs",
		Long:  "Show confirmation dialogs for user decisions",
		Run: func(cmd *cobra.Command, args []string) {
			if message == "" {
				fmt.Fprintf(os.Stderr, "Error: --message is required\n")
				os.Exit(1)
			}

			result, err := tui.EnhancedShowConfirm(message, defaultVal)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			fmt.Println(result)
		},
	}

	cmd.Flags().StringVar(&message, "message", "", "Confirmation message (required)")
	cmd.Flags().StringVar(&defaultVal, "default", "", "Default value (yes/no)")

	return cmd
}
