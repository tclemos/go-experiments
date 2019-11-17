package main

import (
	"context"
	"flag"
	"math"
	"net"
	"os"
	"sync"
	"time"

	"github.com/tclemos/go-experiments/grpc/protocol"
	"github.com/tclemos/go-experiments/grpc/sum"

	"google.golang.org/grpc"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	configureLog()
	serverAddress := serve()

	cnn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create connection")
	}
	defer cnn.Close()

	client := protocol.NewSumServiceClient(cnn)

	ctx := context.Background()
	req := &protocol.SumRequest{
		A: 1,
		B: 1,
	}

	cycles := 10
	n := 100000
	log.Info().Msgf("Starting %d cycles of %d request", cycles, n)
	intervals := []float64{}
	totalStart := time.Now()
	for x := 0; x <= cycles; x++ {
		log.Info().Msgf("Starting cycle %d", x)
		start := time.Now()
		wg := sync.WaitGroup{}

		for i := 0; i < n; i++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup, client protocol.SumServiceClient) {
				defer wg.Done()
				res, err := client.Sum(ctx, req)
				if err != nil {
					log.Error().Err(err).Interface("req", req).Msg("Failed to SUM")
				}

				if res.Sum != (req.A + req.B) {
					log.Error().Err(err).Interface("req", req).Interface("res", res).Msg("Wrong SUM result")
				}

			}(&wg, client)
		}
		wg.Wait()
		interval := time.Now().Sub(start).Seconds()
		intervals = append(intervals, interval)
		log.Info().Msgf("Cycle %d finished and took %f seconds", x, interval)
	}
	totalInterval := time.Now().Sub(totalStart).Seconds()

	intervalsSum := float64(0)
	intervalsMin := float64(math.MaxFloat64)
	intervalsMax := float64(0)
	for _, v := range intervals {
		intervalsSum += v
		if intervalsMin > v {
			intervalsMin = v
		}
		if intervalsMax < v {
			intervalsMax = v
		}
	}

	intervalsAvg := intervalsSum / float64((len(intervals)))

	log.Info().Msgf("Executed %d cycles of %d requests in %f seconds", cycles, n, totalInterval)
	log.Info().Msgf("MIN: %f seconds", intervalsMin)
	log.Info().Msgf("AVG: %f seconds", intervalsAvg)
	log.Info().Msgf("MAX: %f seconds", intervalsMax)
}

func configureLog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	debug := flag.Bool("debug", false, "sets log level to debug")

	flag.Parse()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func serve() string {
	serverAddress := "localhost:7777"
	lis, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatal().Msg("Failed to create tcp listener")
	}

	server := grpc.NewServer()
	sum := sum.NewService()

	protocol.RegisterSumServiceServer(server, sum)

	go server.Serve(lis)

	return serverAddress
}
