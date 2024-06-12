package registry

import (
	"fmt"
	"your_project/interfaces"
)

var (
	inputPluginRegistry  = map[string]func() interfaces.InputPlugin{}
	outputPluginRegistry = map[string]func() interfaces.OutputPlugin{}
)

// RegisterInputPlugin registers an input plugin by name.
func RegisterInputPlugin(name string, constructor func() interfaces.InputPlugin) {
	inputPluginRegistry[name] = constructor
}

// GetInputPlugin returns the input plugin by name.
func GetInputPlugin(name string) (interfaces.InputPlugin, error) {
	constructor, ok := inputPluginRegistry[name]
	if !ok {
		return nil, fmt.Errorf("invalid input plugin: %s", name)
	}
	return constructor(), nil
}

// RegisterOutputPlugin registers an output plugin by name.
func RegisterOutputPlugin(name string, constructor func() interfaces.OutputPlugin) {
	outputPluginRegistry[name] = constructor
}

// GetOutputPlugin returns the output plugin by name.
func GetOutputPlugin(name string) (interfaces.OutputPlugin, error) {
	constructor, ok := outputPluginRegistry[name]
	if !ok {
		return nil, fmt.Errorf("invalid output plugin: %s", name)
	}
	return constructor(), nil
}
