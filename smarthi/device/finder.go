package device

// Finder knows how to find an specific device in the network
type Finder interface {

	// Finds devices in the network
	Find() (*Device, error)
}
