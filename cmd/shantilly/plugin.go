package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/helton-godoy/shantilly-cli/pkg/plugins"
	"github.com/spf13/cobra"
)

func createPluginCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "plugin",
		Short: "Plugin management",
		Long:  "Manage Shantilly-CLI plugins for extended functionality",
	}

	rootCmd.AddCommand(createPluginListCmd())
	rootCmd.AddCommand(createPluginInfoCmd())
	rootCmd.AddCommand(createPluginRunCmd())
	rootCmd.AddCommand(createPluginEnableCmd())
	rootCmd.AddCommand(createPluginDisableCmd())
	rootCmd.AddCommand(createPluginCreateCmd())

	return rootCmd
}

func createPluginListCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "List all loaded plugins",
		Long:  "Show information about all loaded plugins",
		Run: func(cmd *cobra.Command, args []string) {
			config := getDefaultPluginConfig()
			manager := plugins.NewPluginManager(config)

			err := manager.LoadPlugins()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error loading plugins: %v\n", err)
				os.Exit(1)
			}

			pluginList := manager.ListPlugins()
			if len(pluginList) == 0 {
				fmt.Println("No plugins loaded")
				return
			}

			fmt.Println("Loaded plugins:")
			for _, info := range pluginList {
				fmt.Printf("  • %s (v%s)\n", info.Name, info.Version)
				fmt.Printf("    %s\n", info.Description)
				if len(info.Commands) > 0 {
					fmt.Printf("    Commands: ")
					for i, cmd := range info.Commands {
						if i > 0 {
							fmt.Printf(", ")
						}
						fmt.Printf("%s", cmd.Name)
					}
					fmt.Println()
				}
				fmt.Println()
			}
		},
	}

	return cmd
}

func createPluginInfoCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "info [plugin-name]",
		Short: "Show detailed plugin information",
		Long:  "Display detailed information about a specific plugin",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			pluginName := args[0]
			config := getDefaultPluginConfig()
			manager := plugins.NewPluginManager(config)

			err := manager.LoadPlugins()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error loading plugins: %v\n", err)
				os.Exit(1)
			}

			plugin, exists := manager.GetPlugin(pluginName)
			if !exists {
				fmt.Fprintf(os.Stderr, "Plugin not found: %s\n", pluginName)
				os.Exit(1)
			}

			fmt.Printf("Plugin: %s\n", plugin.Name())
			fmt.Printf("Version: %s\n", plugin.Version())
			fmt.Printf("Description: %s\n", plugin.Description())

			commands := plugin.Commands()
			if len(commands) > 0 {
				fmt.Println("\nCommands:")
				for _, cmd := range commands {
					fmt.Printf("  • %s\n", cmd.Name)
					fmt.Printf("    %s\n", cmd.Description)
					fmt.Printf("    Usage: %s\n", cmd.Usage)
					fmt.Println()
				}
			}
		},
	}

	return cmd
}

func createPluginRunCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "run [plugin-name] -- [args...]",
		Short: "Execute a plugin",
		Long:  "Run a plugin with the specified arguments",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			pluginName := args[0]
			var pluginArgs []string
			if len(args) > 1 {
				pluginArgs = args[1:]
			}

			config := getDefaultPluginConfig()
			manager := plugins.NewPluginManager(config)

			err := manager.LoadPlugins()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error loading plugins: %v\n", err)
				os.Exit(1)
			}

			err = manager.ExecutePlugin(pluginName, "", pluginArgs)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error executing plugin: %v\n", err)
				os.Exit(1)
			}
		},
	}

	return cmd
}

func createPluginEnableCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "enable [plugin-name]",
		Short: "Enable a plugin",
		Long:  "Add a plugin to the enabled list",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			pluginName := args[0]
			config := getDefaultPluginConfig()
			manager := plugins.NewPluginManager(config)

			// Add to enabled list
			config.Enabled = append(config.Enabled, pluginName)

			// Save config
			configPath := getPluginConfigPath()
			err := manager.SaveConfig(configPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error saving config: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("Plugin enabled: %s\n", pluginName)
		},
	}

	return cmd
}

func createPluginDisableCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "disable [plugin-name]",
		Short: "Disable a plugin",
		Long:  "Add a plugin to the disabled list",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			pluginName := args[0]
			config := getDefaultPluginConfig()
			manager := plugins.NewPluginManager(config)

			// Add to disabled list
			config.Disabled = append(config.Disabled, pluginName)

			// Save config
			configPath := getPluginConfigPath()
			err := manager.SaveConfig(configPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error saving config: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("Plugin disabled: %s\n", pluginName)
		},
	}

	return cmd
}

func createPluginCreateCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create [plugin-name]",
		Short: "Create a new plugin template",
		Long:  "Generate a template for creating a new plugin",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			pluginName := args[0]

			// Create plugin directory
			pluginDir := filepath.Join(".", pluginName)
			err := os.MkdirAll(pluginDir, 0755)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error creating plugin directory: %v\n", err)
				os.Exit(1)
			}

			// Generate plugin template
			template := fmt.Sprintf(`package main

import (
	"fmt"
)

type %sPlugin struct{}

func (p *%sPlugin) Name() string {
	return "%s"
}

func (p *%sPlugin) Version() string {
	return "1.0.0"
}

func (p *%sPlugin) Description() string {
	return "A sample plugin for Shantilly-CLI"
}

func (p *%sPlugin) Execute(args []string) error {
	fmt.Println("Hello from %s plugin!")
	return nil
}

func (p *%sPlugin) Commands() []plugins.PluginCommand {
	return []plugins.PluginCommand{
		{
			Name:        "hello",
			Description: "Say hello",
			Usage:        "hello",
		},
	}
}

// NewPlugin creates a new plugin instance
func NewPlugin() plugins.Plugin {
	return &%sPlugin{}
}
`, pluginName, pluginName, pluginName, pluginName, pluginName, pluginName, pluginName, pluginName)

			pluginFile := filepath.Join(pluginDir, "main.go")
			err = os.WriteFile(pluginFile, []byte(template), 0644)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error creating plugin file: %v\n", err)
				os.Exit(1)
			}

			// Generate go.mod
			goMod := fmt.Sprintf(`module %s

go 1.21

require github.com/helton-godoy/shantilly-cli v0.0.0
`, pluginName)

			goModFile := filepath.Join(pluginDir, "go.mod")
			err = os.WriteFile(goModFile, []byte(goMod), 0644)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error creating go.mod: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("Plugin template created: %s\n", pluginDir)
			fmt.Printf("To build: cd %s && go build -buildmode=plugin -o %s.so\n", pluginDir, pluginName)
		},
	}

	return cmd
}

func getDefaultPluginConfig() plugins.PluginConfig {
	return plugins.PluginConfig{
		PluginDir: "",
		Enabled:   []string{},
		Disabled:  []string{},
	}
}

func getPluginConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".shantilly", "plugins.json")
}
