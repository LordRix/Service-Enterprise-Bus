package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Config holds the configuration for the plugins and related parameters.
type Config struct {
	InputPlugins  []string
	OutputPlugins []string
	PluginConfig  map[string]map[string]string
}

// LoadConfig reads the configuration from the environment variables.
func LoadConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	config := Config{
		InputPlugins:  getEnvAsSlice("INPUT_PLUGINS", []string{}),
		OutputPlugins: getEnvAsSlice("OUTPUT_PLUGINS", []string{}),
		PluginConfig:  loadPluginConfig(),
	}

	return config, nil
}

// Helper function to get environment variable as string with default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// Helper function to get environment variable as slice with default value
func getEnvAsSlice(key string, defaultValue []string) []string {
	if value, exists := os.LookupEnv(key); exists {
		return strings.Split(value, ",")
	}
	return defaultValue
}

// Load plugin-specific configurations dynamically.
func loadPluginConfig() map[string]map[string]string {
	pluginConfig := make(map[string]map[string]string)
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		key, value := pair[0], pair[1]
		parts := strings.SplitN(key, "_", 3)
		if len(parts) == 3 {
			pluginType, pluginName, configKey := parts[0], parts[1], parts[2]
			if pluginType == "INPUT" || pluginType == "OUTPUT" {
				if _, exists := pluginConfig[pluginName]; !exists {
					pluginConfig[pluginName] = make(map[string]string)
				}
				pluginConfig[pluginName][configKey] = value
			}
		}
	}
	return pluginConfig
}
