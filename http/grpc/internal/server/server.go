package server

import (
	"net"

	"github.com/rs/zerolog/log"
	"github.com/tclemos/go-experiments/http/grpc/protocol"
	"github.com/tclemos/go-experiments/http/grpc/sum"
	"google.golang.org/grpc"
)

func Start(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal().Msg("Failed to create tcp listener")
	}

	server := grpc.NewServer()
	sum := sum.NewService()

	protocol.RegisterSumServiceServer(server, sum)

	server.Serve(lis)
}
