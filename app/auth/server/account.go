package server

import (
	"context"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/authpb"
	"github.com/kimsehyoung/dongle/app/auth/utils/validator"
	"github.com/kimsehyoung/logger"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *authService) CreateAccount(ctx context.Context, req *authpb.AccountRequest) (*emptypb.Empty, error) {
	logger.Info(req.Email, req.Name, req.Password, req.PhoneNumber, req.Role)

	err := validator.IsValidAccount(&validator.Account{
		Email:       req.Email,
		Name:        req.Name,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
	})

	if err != nil {
		logger.Info("Invalid account: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "Failed to validate Account: %v ", err)
	}

	//  This does not accept passwords longer than 72 bytes.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Failed to generate hashed password: %v", err)
	}

	account, err := s.authdbClient.Account.
		Create().
		SetEmail(req.Email).
		SetName(req.Name).
		SetHashedPassword(string(hashedPassword)).
		SetPhoneNumber(req.PhoneNumber).
		SetRoleID(int32(req.Role)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	logger.Info("The account %s(%d) has been created.", account.Email, account.RoleID)

	return nil, nil
}

func (s *authService) UpdateAccount(ctx context.Context, req *authpb.AccountRequest) (*emptypb.Empty, error) {
	logger.Info(req.Email, req.Name, req.Password, req.PhoneNumber, req.Role)
	return nil, nil
}

func (s *authService) DeleteAccount(ctx context.Context, req *authpb.AccountRequest) (*emptypb.Empty, error) {
	logger.Info(req.Email, req.Name, req.Password, req.PhoneNumber, req.Role)
	return nil, nil
}
