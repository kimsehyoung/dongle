package auth

import (
	"context"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/authpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AuthService) CreateAccount(ctx context.Context, req *authpb.AccountRequest) (*emptypb.Empty, error) {

	return nil, nil
}

func (s *AuthService) UpdateAccount(ctx context.Context, req *authpb.AccountRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *AuthService) DeleteAccount(ctx context.Context, req *authpb.AccountRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *AuthService) CreateToken(ctx context.Context, req *authpb.CreateTokenRequest) (*authpb.TokenResposne, error) {
	return nil, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, req *authpb.RefreshTokenRequest) (*authpb.TokenResposne, error) {
	return nil, nil
}

func (s *AuthService) RevokeToken(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, nil
}
