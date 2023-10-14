package server

import (
	"context"
	"encoding/hex"
	"strconv"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/authpb"
	"github.com/kimsehyoung/dongle/app/auth/ent/authdbgen"
	"github.com/kimsehyoung/dongle/app/auth/ent/authdbgen/account"
	"github.com/kimsehyoung/dongle/app/auth/server/jwt"
	"github.com/kimsehyoung/dongle/internal/utils/encrypt"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *authService) CreateToken(ctx context.Context, req *authpb.CreateTokenRequest) (*authpb.TokenResposne, error) {
	// Query account by login ID
	account, err := s.authdbClient.Account.
		Query().
		Where(account.Email(req.Email)).
		Only(ctx)
	if authdbgen.IsNotFound(err) {
		return nil, status.Errorf(codes.InvalidArgument, "Incorrect ID or Password: %v", err)
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to query: %v", err)
	}

	// Compare hashed password
	err = bcrypt.CompareHashAndPassword([]byte(account.HashedPassword), []byte(req.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, status.Errorf(codes.InvalidArgument, "Incorrect ID or Password: %v", err)
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to compare password: %v", err)
	}

	// Create jwt
	accessToken, err := jwt.CreateJwt(account.RoleID, strconv.Itoa(account.ID), jwt.ACCESS_TOKEN)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create access token (%v)", err)
	}
	refreshToken, err := jwt.CreateJwt(account.RoleID, strconv.Itoa(account.ID), jwt.REFRESH_TOKEN)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create refresh token (%v)", err)
	}

	// Response
	return &authpb.TokenResposne{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (s *authService) RefreshToken(ctx context.Context, req *authpb.TokenRequest) (*authpb.TokenResposne, error) {
	// Check whether the refresh toekn is valid or not.
	claims, err := jwt.VerifyJwt(req.RefreshToken)
	if err != nil {
		return nil, err
	}

	plainSubject, err := encrypt.Decrypt([]byte(claims.Subject))
	if err != nil {
		return nil, err
	}

	// Create jwt
	accessToken, err := jwt.CreateJwt(claims.RoleId, hex.EncodeToString(plainSubject), jwt.ACCESS_TOKEN)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create access token (%v)", err)
	}
	refreshToken, err := jwt.CreateJwt(claims.RoleId, hex.EncodeToString(plainSubject), jwt.REFRESH_TOKEN)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create refresh token (%v)", err)
	}

	// Response
	return &authpb.TokenResposne{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (s *authService) RevokeToken(ctx context.Context, req *authpb.TokenRequest) (*emptypb.Empty, error) {
	return nil, nil
}
