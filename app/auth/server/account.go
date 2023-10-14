package server

import (
	"context"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/authpb"
	"github.com/kimsehyoung/dongle/app/auth/ent/authdbgen"
	"github.com/kimsehyoung/dongle/app/auth/ent/authdbgen/role"
	"github.com/kimsehyoung/dongle/app/auth/server/validator"
	"github.com/kimsehyoung/logger"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *authService) CreateAccount(ctx context.Context, req *authpb.AccountRequest) (*emptypb.Empty, error) {

	// Validate account input values
	err := validator.IsValidAccount(&validator.Account{
		Email:       req.Email,
		Name:        req.Name,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Failed to validate Account: %v ", err)
	}

	//  Generate hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to generate hashed password: %v", err)
	}

	// TODO: Creating 'admin' account with secret.
	if req.Role == "root" || req.Role == "admin" {
		return nil, status.Error(codes.Unauthenticated, "Creating admin account is not permitted.")
	}

	// Query role record from request
	role, err := s.authdbClient.Role.
		Query().
		Where(role.Type(req.Role)).
		Only(ctx)
	if authdbgen.IsNotFound(err) {
		return nil, status.Errorf(codes.InvalidArgument, "Not found role: %v", err)
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create: %v", err)
	}

	// Create Acccout
	account, err := s.authdbClient.Account.
		Create().
		SetRole(role).
		SetEmail(req.Email).
		SetHashedPassword(string(hashedPassword)).
		SetName(req.Name).
		SetPhoneNumber(req.PhoneNumber).
		Save(ctx)
	if authdbgen.IsConstraintError(err) {
		return nil, status.Errorf(codes.AlreadyExists, "%s already exists. errmsg(%s)", account.Email, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create: %v", err)
	}

	logger.Infof("The account %s has been created", account.Email)
	return &emptypb.Empty{}, nil
}

func (s *authService) DeleteAccount(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *authService) ChangePassword(ctx context.Context, req *authpb.PasswordRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *authService) ConfirmPassword(ctx context.Context, req *authpb.PasswordRequest) (*emptypb.Empty, error) {
	return nil, nil
}
