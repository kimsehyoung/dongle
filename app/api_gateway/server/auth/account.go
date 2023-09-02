package auth

import (
	"context"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/authpb"
	"github.com/kimsehyoung/logger"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AuthService) CreateAccount(ctx context.Context, req *authpb.AccountRequest) (*emptypb.Empty, error) {
	resp, err := s.AuthClient.CreateAccount(ctx, req)
	if err != nil {
		logger.Errorf("Failed to create account (%v)", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) DeleteAccount(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	resp, err := s.AuthClient.DeleteAccount(ctx, req)
	if err != nil {
		logger.Errorf("Failed to delete account (%v)", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) ChangePassword(ctx context.Context, req *authpb.PasswordRequest) (*emptypb.Empty, error) {
	resp, err := s.AuthClient.ChangePassword(ctx, req)
	if err != nil {
		logger.Errorf("Failed to change password (%v)", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) ConfirmPassword(ctx context.Context, req *authpb.PasswordRequest) (*emptypb.Empty, error) {
	resp, err := s.AuthClient.ConfirmPassword(ctx, req)
	if err != nil {
		logger.Errorf("Failed to confirm password (%v)", err)
		return nil, err
	}
	return resp, nil
}
