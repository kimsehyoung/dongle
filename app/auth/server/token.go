package server

import (
	"context"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/authpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *authService) CreateToken(ctx context.Context, req *authpb.CreateTokenRequest) (*authpb.TokenResposne, error) {
	return nil, nil
}

func (s *authService) RefreshToken(ctx context.Context, req *authpb.TokenRequest) (*authpb.TokenResposne, error) {
	return nil, nil
}

func (s *authService) RevokeToken(ctx context.Context, req *authpb.TokenRequest) (*emptypb.Empty, error) {
	return nil, nil
}
