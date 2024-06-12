package interfaces

// OutputPlugin defines the interface for output plugins.
type OutputPlugin interface {
	SendOutput(output string) error
	Configure(config map[string]string)
}
