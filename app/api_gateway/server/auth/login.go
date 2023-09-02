package auth

import (
	"context"
	"time"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/authpb"
	"github.com/kimsehyoung/dongle/app/auth/server/jwt"
	"github.com/kimsehyoung/logger"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	REFRES_TOKEN_EXPIRATION = 1 * 24 * time.Hour
)

func (s *AuthService) CreateToken(ctx context.Context, req *authpb.CreateTokenRequest) (*authpb.TokenResposne, error) {
	// Request to check whether user info is correct or not.
	resp, err := s.AuthClient.CreateToken(ctx, req)
	if err != nil {
		logger.Errorf("CreateToken (%v)", err)
		return nil, err
	}

	if err = s.TokenDbClient.Set(ctx, resp.RefreshToken, req.Email, jwt.RefreshTokenDuration).Err(); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to store token: %v", err)
	}

	return resp, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, req *authpb.TokenRequest) (*authpb.TokenResposne, error) {

	email, err := s.TokenDbClient.GetDel(ctx, req.RefreshToken).Result()
	if err == redis.Nil {
		return nil, status.Errorf(codes.Unauthenticated, "Failed to store token: %v", err)
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to store token: %v", err)
	}

	resp, err := s.AuthClient.RefreshToken(ctx, req)
	if err != nil {
		logger.Errorf("RefreshToken (%v)", err)
		return nil, err
	}

	if err = s.TokenDbClient.Set(ctx, resp.RefreshToken, email, jwt.RefreshTokenDuration).Err(); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to store token: %v", err)
	}

	return resp, nil
}

func (s *AuthService) RevokeToken(ctx context.Context, req *authpb.TokenRequest) (*emptypb.Empty, error) {

	if err := s.TokenDbClient.Del(ctx, req.RefreshToken).Err(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Failed to delete token: %v", err)
	}

	return &emptypb.Empty{}, nil
}
