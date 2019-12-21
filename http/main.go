package main

import (
	"github.com/tclemos/go-experiments/http/config"
	"github.com/tclemos/go-experiments/http/grpc"
)

func main() {
	config.SetupLog()
	grpc.Run()
	chi.Run()
}
