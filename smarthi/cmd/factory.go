package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tclemos/go-experiments/smarthi/device"
)

// Factory help commands to be created
type Factory interface {

	// CreatesRoot creates the root command
	NewRoot() *cobra.Command

	// CreateToggle creates a command to toggle device properties
	NewToggle() *cobra.Command
}

// BasicFactory implements Factory
type BasicFactory struct {
	deviceFinders []device.Finder
}

// NewBasicFactory Creates and initializes an instance of BasicFactory
func NewBasicFactory(dfs []device.Finder) Factory {
	return &BasicFactory{
		deviceFinders: dfs,
	}
}

// NewRoot creates a new root command
func (f *BasicFactory) NewRoot() *cobra.Command {
	return newRoot()
}

// NewToggle creates a new toggle command
func (f *BasicFactory) NewToggle() *cobra.Command {
	return newToggle(f.deviceFinders)
}
