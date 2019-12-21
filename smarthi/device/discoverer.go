package device

// Discoverer knows how identify devices in the network and provides a list of the devices found
type Discoverer interface {

	// Discover devices in the network
	Discover() ([]Device, error)
}
