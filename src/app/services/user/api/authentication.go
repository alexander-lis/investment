package api

import (
	"context"
	"fmt"

	"github.com/alexander-lis/investment/src/app/shared/protobuf/user/authentication"
)

type AuthentiationServiceServerImpl struct {
	authentication.UnimplementedAuthenticationServiceServer
}

func (o *AuthentiationServiceServerImpl) SignUp(ctx context.Context, in *authentication.SignUpRequest) (*authentication.SignUpResponse, error) {
	fmt.Printf("Signed up")
	return &authentication.SignUpResponse{}, nil
}

func (o *AuthentiationServiceServerImpl) SignIn(ctx context.Context, in *authentication.SignInRequest) (*authentication.SignInResponse, error) {
	fmt.Printf("Signed in")
	return &authentication.SignInResponse{}, nil
}

func (o *AuthentiationServiceServerImpl) LogOut(ctx context.Context, in *authentication.LogOutRequest) (*authentication.LogOutResponse, error) {
	fmt.Printf("Logged out")
	return &authentication.LogOutResponse{}, nil
}
