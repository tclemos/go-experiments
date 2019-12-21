package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tclemos/go-experiments/smarthi/device"
)

// newToggle creates and initializes an instance of Command
func newToggle(dfs []device.Finder) *cobra.Command {
	c := &cobra.Command{
		Use:   "toggle",
		Short: "toggles a property of a device",
		Long: `toggle should be used when a device has a property that can be toggled,
the current property state is not important and you want its value to be toggled.`,
		Args: nil,
		Run: func(cmd *cobra.Command, args []string) {
			run(cmd, args, dfs)
		},
	}

	c.Flags().String("name", "", "device name to have the property toggled")
	c.Flags().String("prop", "", "property name to be toggled")

	return c
}

func run(cmd *cobra.Command, args []string, dfs []device.Finder) {
	fmt.Printf("Toggling the %s %s\n", cmd.Flag("name").Value, cmd.Flag("prop").Value)
}
