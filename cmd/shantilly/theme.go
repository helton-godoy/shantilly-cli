package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/helton-godoy/shantilly-cli/pkg/tui"
	"github.com/spf13/cobra"
)

func createThemeCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "theme",
		Short: "Theme management",
		Long:  "Manage visual themes for Shantilly-CLI TUI components",
	}

	rootCmd.AddCommand(createThemeListCmd())
	rootCmd.AddCommand(createThemeSetCmd())
	rootCmd.AddCommand(createThemeShowCmd())
	rootCmd.AddCommand(createThemeDemoCmd())

	return rootCmd
}

func createThemeListCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "List available themes",
		Long:  "Show all available visual themes",
		Run: func(cmd *cobra.Command, args []string) {
			themes := tui.GetThemeNames()
			fmt.Println("Available themes:")
			for _, theme := range themes {
				fmt.Printf("  • %s\n", theme)
			}
		},
	}

	return cmd
}

func createThemeSetCmd() *cobra.Command {
	var configPath string

	var cmd = &cobra.Command{
		Use:   "set [theme-name]",
		Short: "Set current theme",
		Long:  "Set the default theme for TUI components",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			themeName := args[0]

			// Validate theme
			validThemes := tui.GetThemeNames()
			isValid := false
			for _, theme := range validThemes {
				if theme == themeName {
					isValid = true
					break
				}
			}

			if !isValid {
				fmt.Fprintf(os.Stderr, "Invalid theme: %s\n", themeName)
				fmt.Fprintf(os.Stderr, "Available themes: %s\n", strings.Join(validThemes, ", "))
				os.Exit(1)
			}

			// Load current config
			config, err := tui.LoadConfig(configPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
				os.Exit(1)
			}

			// Update theme
			config.Theme = themeName

			// Save config
			if configPath == "" {
				configPath, err = tui.GetConfigPath()
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error getting config path: %v\n", err)
					os.Exit(1)
				}
			}

			err = tui.SaveConfig(config, configPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error saving config: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("Theme set to: %s\n", themeName)
		},
	}

	cmd.Flags().StringVar(&configPath, "config", "", "Config file path")

	return cmd
}

func createThemeShowCmd() *cobra.Command {
	var configPath string

	var cmd = &cobra.Command{
		Use:   "show",
		Short: "Show current theme",
		Long:  "Display the currently configured theme",
		Run: func(cmd *cobra.Command, args []string) {
			config, err := tui.LoadConfig(configPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("Current theme: %s\n", config.Theme)

			// Show theme details
			theme := tui.GetTheme(config.Theme)
			fmt.Printf("Theme details: %s\n", theme.Name)
		},
	}

	cmd.Flags().StringVar(&configPath, "config", "", "Config file path")

	return cmd
}

func createThemeDemoCmd() *cobra.Command {
	var themeName string

	var cmd = &cobra.Command{
		Use:   "demo",
		Short: "Demo theme with dialog",
		Long:  "Show a demo dialog using the specified theme",
		Run: func(cmd *cobra.Command, args []string) {
			if themeName == "" {
				// Load from config
				config, err := tui.LoadConfig("")
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
					os.Exit(1)
				}
				themeName = config.Theme
			}

			// Demo dialog
			choices := []string{"Option Alpha", "Option Beta", "Option Gamma", "Option Delta"}
			result, err := tui.ThemedShowDialog("Theme Demo", choices, "Option Alpha", false, themeName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("Selected: %s\n", result[0])
		},
	}

	cmd.Flags().StringVar(&themeName, "theme", "", "Theme to demo (default: current config)")

	return cmd
}
