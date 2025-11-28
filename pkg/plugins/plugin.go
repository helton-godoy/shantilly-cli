package plugins

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"plugin"
)

// Plugin interface defines the contract for all plugins
type Plugin interface {
	Name() string
	Version() string
	Description() string
	Execute(args []string) error
	Commands() []PluginCommand
}

// PluginCommand represents a command provided by a plugin
type PluginCommand struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Usage       string `json:"usage"`
}

// PluginInfo contains metadata about a plugin
type PluginInfo struct {
	Name        string          `json:"name"`
	Version     string          `json:"version"`
	Description string          `json:"description"`
	Author      string          `json:"author"`
	Commands    []PluginCommand `json:"commands"`
	Enabled     bool            `json:"enabled"`
}

// PluginManager manages loading and execution of plugins
type PluginManager struct {
	plugins map[string]Plugin
	config  PluginConfig
}

// PluginConfig contains plugin manager configuration
type PluginConfig struct {
	PluginDir string   `json:"pluginDir"`
	Enabled   []string `json:"enabled"`
	Disabled  []string `json:"disabled"`
}

// NewPluginManager creates a new plugin manager
func NewPluginManager(config PluginConfig) *PluginManager {
	return &PluginManager{
		plugins: make(map[string]Plugin),
		config:  config,
	}
}

// LoadPlugins loads all plugins from the plugin directory
func (pm *PluginManager) LoadPlugins() error {
	if pm.config.PluginDir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to get home directory: %w", err)
		}
		pm.config.PluginDir = filepath.Join(home, ".shantilly", "plugins")
	}

	// Create plugin directory if it doesn't exist
	if err := os.MkdirAll(pm.config.PluginDir, 0755); err != nil {
		return fmt.Errorf("failed to create plugin directory: %w", err)
	}

	// Load .so files
	files, err := filepath.Glob(filepath.Join(pm.config.PluginDir, "*.so"))
	if err != nil {
		return fmt.Errorf("failed to glob plugin files: %w", err)
	}

	for _, file := range files {
		if err := pm.loadPlugin(file); err != nil {
			fmt.Printf("Warning: failed to load plugin %s: %v\n", file, err)
		}
	}

	return nil
}

// loadPlugin loads a single plugin from file
func (pm *PluginManager) loadPlugin(file string) error {
	p, err := plugin.Open(file)
	if err != nil {
		return fmt.Errorf("failed to open plugin: %w", err)
	}

	// Look for NewPlugin function
	newPluginSymbol, err := p.Lookup("NewPlugin")
	if err != nil {
		return fmt.Errorf("plugin does not export NewPlugin: %w", err)
	}

	// Type assert to function
	newPlugin, ok := newPluginSymbol.(func() Plugin)
	if !ok {
		return fmt.Errorf("NewPlugin has wrong signature")
	}

	// Create plugin instance
	pluginInstance := newPlugin()

	// Check if plugin is enabled
	if !pm.isPluginEnabled(pluginInstance.Name()) {
		return nil
	}

	// Register plugin
	pm.plugins[pluginInstance.Name()] = pluginInstance

	return nil
}

// isPluginEnabled checks if a plugin is enabled
func (pm *PluginManager) isPluginEnabled(name string) bool {
	// Check disabled list
	for _, disabled := range pm.config.Disabled {
		if disabled == name {
			return false
		}
	}

	// If enabled list is empty, all plugins are enabled by default
	if len(pm.config.Enabled) == 0 {
		return true
	}

	// Check enabled list
	for _, enabled := range pm.config.Enabled {
		if enabled == name {
			return true
		}
	}

	return false
}

// GetPlugin returns a plugin by name
func (pm *PluginManager) GetPlugin(name string) (Plugin, bool) {
	plugin, exists := pm.plugins[name]
	return plugin, exists
}

// ListPlugins returns information about all loaded plugins
func (pm *PluginManager) ListPlugins() []PluginInfo {
	var plugins []PluginInfo

	for _, plugin := range pm.plugins {
		info := PluginInfo{
			Name:        plugin.Name(),
			Version:     plugin.Version(),
			Description: plugin.Description(),
			Commands:    plugin.Commands(),
			Enabled:     true,
		}
		plugins = append(plugins, info)
	}

	return plugins
}

// ExecutePlugin executes a plugin command
func (pm *PluginManager) ExecutePlugin(pluginName, command string, args []string) error {
	plugin, exists := pm.GetPlugin(pluginName)
	if !exists {
		return fmt.Errorf("plugin not found: %s", pluginName)
	}

	// For now, just execute the plugin with all args
	// TODO: Add command-specific execution
	return plugin.Execute(args)
}

// SaveConfig saves the plugin configuration
func (pm *PluginManager) SaveConfig(configPath string) error {
	data, err := json.MarshalIndent(pm.config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	return os.WriteFile(configPath, data, 0644)
}

// LoadConfig loads the plugin configuration
func (pm *PluginManager) LoadConfig(configPath string) error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	return json.Unmarshal(data, &pm.config)
}
