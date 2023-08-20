package server

import (
	"net"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/authpb"
	"github.com/kimsehyoung/dongle/app/auth/ent/authgen"
	"github.com/kimsehyoung/dongle/app/auth/server/authdb"
	"github.com/kimsehyoung/logger"
	"google.golang.org/grpc"
)

type ServiceInfo struct {
	AuthServiceAddr string
	DatabaseInfo    authdb.DatabaseInfo
}

type authService struct {
	authpb.AuthServer
	authdbClient *authgen.Client
}

func StartGrpcServer(serviceInfo ServiceInfo) {
	// Register gRPC Server
	grpcServer := grpc.NewServer()

	authpb.RegisterAuthServer(grpcServer, &authService{
		authdbClient: authdb.GetDbClient(serviceInfo.DatabaseInfo),
	})

	// Start gRPC server
	lis, err := net.Listen("tcp", serviceInfo.AuthServiceAddr)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	logger.Infof("gRPC server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}
