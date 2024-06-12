package main

import (
    "fmt"
    "os"
    "os/signal"
    "strings"
    "syscall"
    "sync"
    "your_project/config"
    "your_project/log"
    "your_project/registry"
    "your_project/interfaces"
    "github.com/rs/zerolog/log"
)

func main() {
    // Initialize the logger
    log.InitLogger()

    // Load the configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to load configuration")
    }

    var wg sync.WaitGroup

    // Initialize input plugins
    inputPlugins := []interfaces.InputPlugin{}
    for _, pluginName := range cfg.InputPlugins {
        inputPlugin, err := registry.GetInputPlugin(pluginName)
        if err != nil {
            log.Fatal().Err(err).Str("plugin", pluginName).Msg("Failed to get input plugin")
        }
        inputPlugin.Configure(cfg.PluginConfig[pluginName])
        inputPlugins = append(inputPlugins, inputPlugin)
    }

    // Initialize output plugins
    outputPlugins := []interfaces.OutputPlugin{}
    for _, pluginName := range cfg.OutputPlugins {
        outputPlugin, err := registry.GetOutputPlugin(pluginName)
        if err != nil {
            log.Fatal().Err(err).Str("plugin", pluginName).Msg("Failed to get output plugin")
        }
        outputPlugin.Configure(cfg.PluginConfig[pluginName])
        outputPlugins = append(outputPlugins, outputPlugin)
    }

    // Handle inputs
    for _, inputPlugin := range inputPlugins {
        wg.Add(1)
        go func(inputPlugin interfaces.InputPlugin) {
            defer wg.Done()
            for {
                input, err := inputPlugin.GetInput()
                if err != nil {
                    log.Error().Err(err).Msg("Failed to get input")
                    continue
                }

                mappedString := MapString(input)
                log.Info().Str("input", input).Str("mapped", mappedString).Msg("Mapped string")
                fmt.Println("Mapped string:", mappedString)

                for _, outputPlugin := range outputPlugins {
                    err := outputPlugin.SendOutput(mappedString)
                    if err != nil {
                        log.Error().Err(err).Msg("Failed to send output")
                    }
                }
            }
        }(inputPlugin)
    }

    // Set up signal handling to gracefully shut down
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
    <-stop

    wg.Wait()
    log.Info().Msg("Shutting down gracefully...")
}

// MapString is a function that performs some transformation on the input string.
func MapString(input string) string {
    // Example transformation: convert the string to uppercase
    return strings.ToUpper(input)
}
