package authentication

import (
	"context"
)

type server struct {
	AuthenticationServer
}

func (s *server) SayHelloAgain(ctx context.Context, in *SignUpRequest) (*SignUpResponse, error) {
	return &SignUpResponse{}, nil
}
