package log

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger() {
	// Set the global logger to output JSON formatted logs to the standard output
	log.Logger = zerolog.New(os.Stdout).With().
		Timestamp().
		Logger()

	// Set zerolog to use UTC time and include the message key
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.TimestampFieldName = "timestamp"
	zerolog.MessageFieldName = "message"
}
