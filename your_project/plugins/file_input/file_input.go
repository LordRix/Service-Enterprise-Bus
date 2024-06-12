package file_input

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
	"your_project/interfaces"
	"your_project/registry"

	"github.com/rs/zerolog/log"
)

// FileInput implements InputPlugin for reading files in a directory.
type FileInput struct {
	directories []string
}

// NewFileInput creates a new FileInput instance.
func NewFileInput() interfaces.InputPlugin {
	return &FileInput{}
}

// Configure sets up the FileInput with the given configuration.
func (fi *FileInput) Configure(config map[string]string) {
	paths := config["PATHS"]
	if paths == "" {
		log.Fatal().Msg("FileInput: PATHS is required")
	}
	fi.directories = strings.Split(paths, ",")
}

// GetInput reads new files from the configured directories.
func (fi *FileInput) GetInput() (string, error) {
	for _, dir := range fi.directories {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			log.Error().Err(err).Str("directory", dir).Msg("Failed to read directory")
			return "", err
		}

		for _, file := range files {
			if file.Mode().IsRegular() {
				filePath := filepath.Join(dir, file.Name())
				content, err := ioutil.ReadFile(filePath)
				if err != nil {
					log.Error().Err(err).Str("file_path", filePath).Msg("Failed to read file")
					continue
				}
				os.Remove(filePath) // simulate processing by deleting file
				log.Info().Str("file_path", filePath).Msg("Processed file and removed")
				return string(content), nil
			}
		}
	}
	time.Sleep(5 * time.Second) // Simulate delay for continuous monitoring
	log.Info().Msg("No new files found; will check again")
	return "", fmt.Errorf("no new files")
}

func init() {
	registry.RegisterInputPlugin("file", NewFileInput)
}
