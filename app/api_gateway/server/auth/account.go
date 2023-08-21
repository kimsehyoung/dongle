package auth

import (
	"context"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/authpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AuthService) CreateAccount(ctx context.Context, req *authpb.AccountRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *AuthService) DeleteAccount(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *AuthService) ChangePassword(ctx context.Context, req *authpb.PasswordRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *AuthService) ConfirmPassword(ctx context.Context, req *authpb.PasswordRequest) (*emptypb.Empty, error) {
	return nil, nil
}
