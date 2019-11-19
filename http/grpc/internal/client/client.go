package client

import (
	"context"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/tclemos/go-experiments/http/config"
	"github.com/tclemos/go-experiments/http/grpc/protocol"
	"google.golang.org/grpc"
)

func Execute(addr string) []time.Duration {
	cnn := createConnection(addr)
	defer cnn.Close()

	client := protocol.NewSumServiceClient(cnn)

	return executeCycles(client)
}

func createConnection(addr string) *grpc.ClientConn {
	cnn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create connection")
	}
	return cnn
}

func executeCycles(client protocol.SumServiceClient) []time.Duration {
	results := []time.Duration{}

	log.Info().Msgf("Starting %d cycles of %d request", config.Cycles, config.Requests)
	for c := 0; c < config.Cycles; c++ {
		start := time.Now()
		executeRequests(c, client)
		interval := time.Now().Sub(start)
		log.Info().Msgf("Executed cycle %d in %f seconds", c, interval.Seconds())
		results = append(results, interval)
	}
	return results
}

func executeRequests(cycle int, client protocol.SumServiceClient) {
	log.Info().Msgf("Starting cycle %d", cycle)
	wg := &sync.WaitGroup{}
	for r := 0; r < config.Requests; r++ {
		wg.Add(1)
		go func(client protocol.SumServiceClient, wg *sync.WaitGroup) {
			defer wg.Done()
			executeRequest(client)
		}(client, wg)
	}
	wg.Wait()
}

func executeRequest(client protocol.SumServiceClient) {
	req := &protocol.SumRequest{
		A: 1,
		B: 1,
	}

	res, err := client.Sum(context.Background(), req)
	if err != nil {
		log.Fatal().Err(err).Interface("req", req).Msg("Failed to SUM")
	}

	if res.Sum != (req.A + req.B) {
		log.Fatal().Err(err).Interface("req", req).Interface("res", res).Msg("Wrong SUM result")
	}
}
