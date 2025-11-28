package main

import (
	"fmt"
)

type samplePlugin struct{}

func (p *samplePlugin) Name() string {
	return "sample"
}

func (p *samplePlugin) Version() string {
	return "1.0.0"
}

func (p *samplePlugin) Description() string {
	return "A sample plugin for Shantilly-CLI"
}

func (p *samplePlugin) Execute(args []string) error {
	fmt.Println("Hello from sample plugin!")
	return nil
}

func (p *samplePlugin) Commands() []plugins.PluginCommand {
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
	return &%!s(MISSING)Plugin{}
}
