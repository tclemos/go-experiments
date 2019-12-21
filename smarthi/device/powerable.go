package device

// Powerable defines a device that can be turned On and Off
type Powerable interface {
	// IsOn returns true when the device is on; otherwise false.
	IsOn() bool

	// On turns it on
	On() error

	// Off turns it off
	Off() error

	// Toggle toggles whether it is on of off
	Toggle() error
}
