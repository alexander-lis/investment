package authentication

import (
	"context"

	"github.com/alexander-lis/money/app/protobuf/user/authentication"
)

type server struct {
	authentication.AuthenticationServer
}

func (s *server) SignUp(ctx context.Context, in *authentication.SignUpRequest) (*authentication.SignUpResponse, error) {
	return &authentication.SignUpResponse{Name: in.Name}, nil
}

func (s *server) SignIn(ctx context.Context, in *authentication.SignInRequest) (*authentication.SignInResponse, error) {
	return &authentication.SignInResponse{}, nil
}

func (s *server) LogOut(ctx context.Context, in *authentication.LogOutRequest) (*authentication.LogOutResponse, error) {
	return &authentication.LogOutResponse{}, nil
}
