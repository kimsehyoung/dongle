package server

import (
	"context"
	"net"
	"net/http"

	auth_interceptor "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/selector"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kimsehyoung/dongle/api/proto/gen/go/authpb"
	"github.com/kimsehyoung/dongle/api/proto/gen/go/speechpb"
	"github.com/kimsehyoung/dongle/app/api_gateway/server/auth"
	"github.com/kimsehyoung/dongle/app/api_gateway/server/speech"
	"github.com/kimsehyoung/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceInfo struct {
	// API Gateway gRPC, REST Ports
	GrpcPort string
	RestPort string
	// TLS configs
	TlsEnabled  bool
	KeyFilePath string
	CrtFilePath string
	// Services
	TestServiceAddr   string
	AuthServiceAddr   string
	SpeechServiceAddr string
}

func StartGrpcServer(serviceInfo ServiceInfo) {

	// Configure gRPC server
	opts := []grpc.ServerOption{
		// https://github.dev/grpc-ecosystem/go-grpc-middleware/blob/main/interceptors/auth/auth.go
		// https://pkg.go.dev/github.com/grpc-ecosystem/go-grpc-middleware/v2#section-readme
		grpc.ChainUnaryInterceptor(
			selector.UnaryServerInterceptor(auth_interceptor.UnaryServerInterceptor(auth.AuthInterceptor), selector.MatchFunc(auth.AuthSkip)),
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

	// Register servicesb
	authpb.RegisterAuthServer(grpcServer, &auth.AuthService{AuthClient: auth.GetAuthClient(serviceInfo.AuthServiceAddr)})
	speechpb.RegisterSpeechServer(grpcServer, &speech.SpeechService{SpeechClient: speech.GetSpeechClient(serviceInfo.SpeechServiceAddr)})

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

	// Start REST server
	logger.Infof("Start gRPC-Gateway server(TLS: %t) listening at %s", serviceInfo.TlsEnabled, serviceInfo.RestPort)
	if serviceInfo.TlsEnabled {
		err = http.ListenAndServeTLS("0.0.0.0:"+serviceInfo.RestPort, serviceInfo.CrtFilePath, serviceInfo.KeyFilePath, mux)
	} else {
		err = http.ListenAndServe("0.0.0.0:"+serviceInfo.RestPort, mux)
	}
	if err != nil {
		logger.Fatalf("runServer error: %v", err)
	}
}
