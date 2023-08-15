package server

import (
	"context"
	"net"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/speechpb"
	"github.com/kimsehyoung/dongle/app/speech/server/google"
	"github.com/kimsehyoung/dongle/app/speech/server/whisper"
	"github.com/kimsehyoung/gopackages/shlog"
	"google.golang.org/grpc"
)

type ServiceInfo struct {
	SpeechServiceAddr string
	WhisperServerAddr string
}

type speechService struct {
	speechpb.SpeechServer
	whisperClient *whisper.WhisperClient
	googleClient  *google.GoogleClient
}

func StartGrpcServer(serviceInfo ServiceInfo) {
	// Register gRPC Server
	grpcServer := grpc.NewServer()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	speechpb.RegisterSpeechServer(grpcServer, &speechService{
		whisperClient: whisper.GetWhisperClient(serviceInfo.WhisperServerAddr),
		googleClient:  google.GetGoogleClient(ctx),
	})

	// Start gRPC server
	lis, err := net.Listen("tcp", serviceInfo.SpeechServiceAddr)
	if err != nil {
		shlog.Logf("FATAL", "failed to listen: %v", err)
	}

	shlog.Logf("INFO", "gRPC server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		shlog.Logf("FATAL", "failed to serve: %v", err)
	}
}
