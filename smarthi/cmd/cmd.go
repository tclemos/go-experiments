package cmd

const rootCommandName = "smarthi"

// Setup all the commands
func Setup(fac Factory) {
	toggle := fac.NewToggle()

	root := fac.NewRoot()

	root.AddCommand(toggle)

	root.Execute()
}
