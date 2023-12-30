package handlers

import (
	"context"

	"github.com/alexander-lis/investment/src/app/shared/protobuf/user/authentication"
)

type AuthentiationServiceServerImpl struct {
}

func (o *AuthentiationServiceServerImpl) SignUp(ctx context.Context, in *authentication.SignUpRequest) (*authentication.SignUpResponse, error) {
	return nil, nil
}

func (o *AuthentiationServiceServerImpl) SignIn(ctx context.Context, in *authentication.SignInRequest) (*authentication.SignInResponse, error) {
	return nil, nil
}

func (o *AuthentiationServiceServerImpl) LogOut(ctx context.Context, in *authentication.LogOutRequest) (*authentication.LogOutResponse, error) {
	return nil, nil
}
