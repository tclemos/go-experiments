package cmd

import "github.com/spf13/cobra"

func newRoot() *cobra.Command {
	return &cobra.Command{
		Use: rootCommandName,
	}
}
