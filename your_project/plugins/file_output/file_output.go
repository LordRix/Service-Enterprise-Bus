package file_output

import (
	"os"
	"your_project/interfaces"
	"your_project/registry"

	"github.com/rs/zerolog/log"
)

// FileOutput implements OutputPlugin for writing to a file.
type FileOutput struct {
	filePath string
}

// NewFileOutput creates a new FileOutput instance.
func NewFileOutput() interfaces.OutputPlugin {
	return &FileOutput{}
}

// Configure sets up the FileOutput with the given configuration.
func (fo *FileOutput) Configure(config map[string]string) {
	fo.filePath = config["FILE_PATH"]
	if fo.filePath == "" {
		log.Fatal().Msg("FileOutput: FILE_PATH is required")
	}
}

// SendOutput writes the output message to the specified file.
func (fo *FileOutput) SendOutput(output string) error {
	file, err := os.OpenFile(fo.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Error().Err(err).Str("file_path", fo.filePath).Msg("Failed to open file for writing")
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(output + "\n"); err != nil {
		log.Error().Err(err).Str("file_path", fo.filePath).Msg("Failed to write to file")
		return err
	}

	log.Info().Str("file_path", fo.filePath).Msg("Output written to file")
	return nil
}

func init() {
	registry.RegisterOutputPlugin("file", NewFileOutput)
}
