package server

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kimsehyoung/dongle/api/proto/gen/go/authpb"
	"github.com/kimsehyoung/dongle/api/proto/gen/go/speechpb"
	"github.com/kimsehyoung/dongle/api/proto/gen/go/videopb"
	"github.com/kimsehyoung/dongle/app/api_gateway/server/auth"
	"github.com/kimsehyoung/dongle/app/api_gateway/server/interceptor"
	"github.com/kimsehyoung/dongle/app/api_gateway/server/speech"
	"github.com/kimsehyoung/dongle/app/api_gateway/server/video"
	"github.com/kimsehyoung/logger"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceInfo struct {
	AllowedOrigins string
	// API Gateway gRPC, REST Ports
	GrpcPort string
	RestPort string
	// TLS configs
	TlsEnabled  bool
	KeyFilePath string
	CrtFilePath string
	// Services
	AuthServiceAddr   string
	SpeechServiceAddr string
	VideoServiceAddr  string
	// Redis for refresh token
	TokenDbAddr string
}

func StartGrpcServer(serviceInfo ServiceInfo) {

	// Configure gRPC server
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			interceptor.AuthUnaryInterceptor(),
		),
	}
	if serviceInfo.TlsEnabled {
		creds, err := credentials.NewServerTLSFromFile(serviceInfo.CrtFilePath, serviceInfo.KeyFilePath)
		if err != nil {
			logger.Fatalf("TLS configure error: %v", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}
	grpcServer := grpc.NewServer(opts...)

	// Register services
	authpb.RegisterAuthServer(grpcServer, &auth.AuthService{
		AuthClient:    auth.GetAuthClient(serviceInfo.AuthServiceAddr),
		TokenDbClient: auth.GetTokenDbClient(serviceInfo.TokenDbAddr),
	})
	speechpb.RegisterSpeechServer(grpcServer, &speech.SpeechService{SpeechClient: speech.GetSpeechClient(serviceInfo.SpeechServiceAddr)})
	videopb.RegisterVideoServer(grpcServer, &video.VideoService{VideoClient: video.GetVideoClient(serviceInfo.VideoServiceAddr)})

	// Start gRPC Server
	lis, err := net.Listen("tcp", "0.0.0.0:"+serviceInfo.GrpcPort)
	if err != nil {
		logger.Fatalf("gRPC Listen error: %v", err)
	}
	go func() {
		logger.Infof("Start gRPC server(TLS: %t) listening at %s", serviceInfo.TlsEnabled, serviceInfo.GrpcPort)
		err := grpcServer.Serve(lis)
		if err != nil {
			logger.Fatalf("gRPC Serve error: %v", err)
		}
	}()
}

func StartRestServer(serviceInfo ServiceInfo) {
	// Register gRPC server Endpoints
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := authpb.RegisterAuthHandlerFromEndpoint(ctx, mux, "127.0.0.1:"+serviceInfo.GrpcPort, opts)
	if err != nil {
		logger.Fatalf("can't register auth service handler: %v", err)
	}
	err = speechpb.RegisterSpeechHandlerFromEndpoint(ctx, mux, "127.0.0.1:"+serviceInfo.GrpcPort, opts)
	if err != nil {
		logger.Fatalf("can't register speech service handler: %v", err)
	}
	err = videopb.RegisterVideoHandlerFromEndpoint(ctx, mux, "127.0.0.1:"+serviceInfo.GrpcPort, opts)
	if err != nil {
		logger.Fatalf("can't register video service handler: %v", err)
	}

	// CORS
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{serviceInfo.AllowedOrigins},
		AllowedMethods: []string{"POST", "GET", "PUT", "DELETE"},
		AllowedHeaders: []string{"Origin", "Accept", "Content-Type", "Authorization"},
	}).Handler(mux)

	// Start REST server
	logger.Infof("Start gRPC-Gateway server(TLS: %t) listening at %s", serviceInfo.TlsEnabled, serviceInfo.RestPort)
	if serviceInfo.TlsEnabled {
		err = http.ListenAndServeTLS("0.0.0.0:"+serviceInfo.RestPort, serviceInfo.CrtFilePath, serviceInfo.KeyFilePath, handler)
	} else {
		err = http.ListenAndServe("0.0.0.0:"+serviceInfo.RestPort, handler)
	}
	if err != nil {
		logger.Fatalf("runServer error: %v", err)
	}
}
