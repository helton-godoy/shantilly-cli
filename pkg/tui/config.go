package tui

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Config represents the application configuration
type Config struct {
	Theme       string `json:"theme"`
	Animation   bool   `json:"animation"`
	Color       bool   `json:"color"`
	KeyboardNav bool   `json:"keyboardNav"`
	AutoSave    bool   `json:"autoSave"`
	Debug       bool   `json:"debug"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() Config {
	return Config{
		Theme:       "default",
		Animation:   true,
		Color:       true,
		KeyboardNav: true,
		AutoSave:    false,
		Debug:       false,
	}
}

// LoadConfig loads configuration from file
func LoadConfig(configPath string) (Config, error) {
	config := DefaultConfig()

	if configPath == "" {
		// Default config path
		home, err := os.UserHomeDir()
		if err != nil {
			return config, fmt.Errorf("failed to get home directory: %w", err)
		}
		configPath = filepath.Join(home, ".shantilly", "config.json")
	}

	// Check if config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Create default config
		err = SaveConfig(config, configPath)
		return config, err
	}

	// Read config file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return config, fmt.Errorf("failed to read config file: %w", err)
	}

	// Parse JSON
	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("failed to parse config file: %w", err)
	}

	return config, nil
}

// SaveConfig saves configuration to file
func SaveConfig(config Config, configPath string) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Marshal to JSON
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	// Write to file
	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

// GetConfigPath returns the default config path
func GetConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}
	return filepath.Join(home, ".shantilly", "config.json"), nil
}
