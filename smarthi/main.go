package main

import (
	"github.com/tclemos/go-experiments/smarthi/cmd"
	"github.com/tclemos/go-experiments/smarthi/device"
)

func main() {
	dfs := []device.Finder{}
	cf := cmd.NewBasicFactory(dfs)
	cmd.Setup(cf)
}
