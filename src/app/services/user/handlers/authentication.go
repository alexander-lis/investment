package handlers

import (
	user "alexander-lis/investment/shared/protobuf/services/user/proto/v1"
	"context"
	"fmt"
)

type AuthenticationServiceServerImpl struct {
	user.UnimplementedAuthenticationServiceServer
}

func (o *AuthenticationServiceServerImpl) SignUp(ctx context.Context, in *user.SignUpRequest) (*user.SignUpResponse, error) {
	fmt.Printf("Signed up")

	return &user.SignUpResponse{
		Name: "vasiliy",
	}, nil
}

func (o *AuthenticationServiceServerImpl) SignIn(ctx context.Context, in *user.SignInRequest) (*user.SignInResponse, error) {
	fmt.Printf("Signed in")
	return &user.SignInResponse{}, nil
}

func (o *AuthenticationServiceServerImpl) LogOut(ctx context.Context, in *user.LogOutRequest) (*user.LogOutResponse, error) {
	fmt.Printf("Logged out")
	return &user.LogOutResponse{}, nil
}
