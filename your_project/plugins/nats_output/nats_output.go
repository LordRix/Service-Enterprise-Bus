package nats_output

import (
	"your_project/interfaces"
	"your_project/registry"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
)

// NatsOutput implements OutputPlugin for publishing messages to NATS.
type NatsOutput struct {
	server  string
	subject string
	conn    *nats.Conn
}

// NewNatsOutput creates a new NatsOutput instance.
func NewNatsOutput() interfaces.OutputPlugin {
	return &NatsOutput{}
}

// Configure sets up the NatsOutput with the given configuration.
func (no *NatsOutput) Configure(config map[string]string) {
	server := config["SERVER"]
	subject := config["SUBJECT"]
	no.server = server
	no.subject = subject

	conn, err := nats.Connect(server)
	if err != nil {
		log.Fatal().Err(err).Str("nats_server", server).Msg("Failed to connect to NATS server")
	}
	no.conn = conn
}

// SendOutput publishes the output message to the NATS server.
func (no *NatsOutput) SendOutput(output string) error {
	err := no.conn.Publish(no.subject, []byte(output))
	if err != nil {
		log.Error().Err(err).Str("subject", no.subject).Msg("Failed to publish message to NATS")
		return err
	}
	log.Info().Str("subject", no.subject).Str("message", output).Msg("Published message to NATS")
	return nil
}

func init() {
	registry.RegisterOutputPlugin("nats", NewNatsOutput)
}
