package device

import "net"

// Device represents a generic device in the network
type Device struct {
	// IP is the device network address
	IP net.IPAddr

	// Name is an alias to a device in the network
	Name string
}
