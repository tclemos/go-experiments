package config

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	// Cycles represents how many cycles will be executed
	Cycles = 10

	// Requests represents how many requests will be executed for each cycle
	Requests = 100000
)

// SetupLog sets the log configuration
func SetupLog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}
