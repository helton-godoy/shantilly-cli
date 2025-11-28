package main

import (
	"fmt"
	"os"

	"github.com/helton-godoy/shantilly-cli/pkg/tui"
	"github.com/spf13/cobra"
)

func createBubbleTeaCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "bubbletea",
		Short: "BubbleTea TUI components",
		Long:  "Advanced TUI components using BubbleTea framework",
	}

	rootCmd.AddCommand(createBubbleTeaDialogCmd())
	rootCmd.AddCommand(createBubbleTeaProgressCmd())

	return rootCmd
}

func createBubbleTeaDialogCmd() *cobra.Command {
	var title string
	var options []string
	var defaultVal string
	var multi bool

	var cmd = &cobra.Command{
		Use:   "dialog",
		Short: "BubbleTea interactive dialog boxes",
		Long:  "Advanced interactive dialog boxes with keyboard navigation",
		Run: func(cmd *cobra.Command, args []string) {
			if title == "" || len(options) == 0 {
				fmt.Fprintf(os.Stderr, "Error: --title and --options are required\n")
				os.Exit(1)
			}

			parsedOptions := tui.ParseOption(options)
			result, err := tui.BubbleTeaShowDialog(title, parsedOptions, defaultVal, multi)
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
	cmd.Flags().StringArrayVar(&options, "options", []string{}, "Options (required)")
	cmd.Flags().StringVar(&defaultVal, "default", "", "Default selection")
	cmd.Flags().BoolVar(&multi, "multi", false, "Allow multiple selections")

	return cmd
}

func createBubbleTeaProgressCmd() *cobra.Command {
	var steps []string
	var current int

	var cmd = &cobra.Command{
		Use:   "progress",
		Short: "BubbleTea progress indicators",
		Long:  "Animated progress bars with real-time updates",
		Run: func(cmd *cobra.Command, args []string) {
			if len(steps) == 0 {
				fmt.Fprintf(os.Stderr, "Error: --steps is required\n")
				os.Exit(1)
			}

			err := tui.BubbleTeaShowProgress(steps, current)
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
