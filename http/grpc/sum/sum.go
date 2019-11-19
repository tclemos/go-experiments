package sum

import (
	"context"

	"github.com/tclemos/go-experiments/http/grpc/protocol"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Sum(ctx context.Context, req *protocol.SumRequest) (*protocol.SumResponse, error) {

	res := &protocol.SumResponse{
		Sum: req.A + req.B,
	}

	return res, nil
}
