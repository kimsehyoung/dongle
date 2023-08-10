package server

import (
	"net"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/testpb"
	"github.com/kimsehyoung/gopackages/shlog"
	"google.golang.org/grpc"
)

type ServiceInfo struct {
	TestServiceAddr string
}

type testService struct {
	testpb.TestServer
}

func StartGrpcServer(serviceInfo ServiceInfo) {
	lis, err := net.Listen("tcp", serviceInfo.TestServiceAddr)
	if err != nil {
		shlog.Logf("FATAL", "failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	testpb.RegisterTestServer(grpcServer, &testService{})

	shlog.Logf("INFO", "gRPC server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		shlog.Logf("FATAL", "failed to serve: %v", err)
	}
}
