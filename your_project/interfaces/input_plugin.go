package interfaces

// InputPlugin defines the interface for input plugins.
type InputPlugin interface {
	GetInput() (string, error)
	Configure(config map[string]string)
}
