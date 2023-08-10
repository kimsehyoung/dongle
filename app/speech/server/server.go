package server

import (
	"net"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/speechpb"
	"github.com/kimsehyoung/gopackages/shlog"

	"google.golang.org/grpc"
)

type ServiceInfo struct {
	SpeechServiceAddr string
}

type speechService struct {
	speechpb.SpeechServer
}

func StartGrpcServer(serviceInfo ServiceInfo) {
	lis, err := net.Listen("tcp", serviceInfo.SpeechServiceAddr)
	if err != nil {
		shlog.Logf("FATAL", "failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	speechpb.RegisterSpeechServer(grpcServer, &speechService{})

	shlog.Logf("INFO", "gRPC server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		shlog.Logf("FATAL", "failed to serve: %v", err)
	}
}
