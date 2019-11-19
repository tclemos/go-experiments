package grpc

import (
	"math"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/tclemos/go-experiments/http/config"
	"github.com/tclemos/go-experiments/http/grpc/internal/client"
	"github.com/tclemos/go-experiments/http/grpc/internal/server"
)

func Run() {
	log.Info().Msg("Begin of gRPC simulation")
	addr := "localhost:7777"
	go server.Start(addr)
	start := time.Now()
	results := client.Execute(addr)
	total := time.Now().Sub(start)
	printResults(total, results)
	log.Info().Msg("End of the gRPC simulation")
}

func printResults(total time.Duration, results []time.Duration) {
	sum := time.Duration(0)
	min := time.Duration(math.MaxInt64)
	max := time.Duration(0)

	for _, v := range results {
		sum += v
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}

	avg := time.Duration(int64(sum) / int64(len(results)))
	log.Info().Msgf("Executed %d cycles of %d requests in %f seconds", config.Cycles, config.Requests, total.Seconds())
	log.Info().Msgf("MIN: %f seconds", min.Seconds())
	log.Info().Msgf("AVG: %f seconds", avg.Seconds())
	log.Info().Msgf("MAX: %f seconds", max.Seconds())
}
