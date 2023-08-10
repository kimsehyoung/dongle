package server

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kimsehyoung/dongle/api/proto/gen/go/testpb"
	"github.com/kimsehyoung/dongle/app/api_gateway/server/test"
	"github.com/kimsehyoung/gopackages/shlog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceInfo struct {
	GrpcPort string
	RestPort string

	TlsEnabled  bool
	KeyFilePath string
	CrtFilePath string

	TestServiceAddr   string
	AuthServiceAddr   string
	SpeechServiceAddr string
}

func StartGrpcServer(serviceInfo ServiceInfo) {
	// Register gRPC-Gateway server
	grpcServer := grpc.NewServer()

	testpb.RegisterTestServer(grpcServer, &test.TestService{TestClient: test.GetTestClient(serviceInfo.TestServiceAddr)})
	// authpb.RegisterAuthServer(grpcServer, &ApiGatewayServer{})
	// speechpb.RegisterSpeechServer(grpcServer, &ApiGatewayServer{})

	// Start gRPC Server
	lis, err := net.Listen("tcp", "0.0.0.0:"+serviceInfo.GrpcPort)
	if err != nil {
		shlog.Logf("FATAL", "gRPC Listen error: %v", err)
	}
	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			shlog.Logf("FATAL", "gRPC Serve error: %v", err)
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
	err := testpb.RegisterTestHandlerFromEndpoint(ctx, mux, "127.0.0.1:"+serviceInfo.GrpcPort, opts)
	if err != nil {
		shlog.Logf("FATAL", "can't register test service handler: %v", err)
	}
	// err = authpb.RegisterAuthHandlerFromEndpoint(ctx, mux, "127.0.0.1:"+serviceInfo.GrpcPort, opts)
	// if err != nil {
	// 	shlog.Logf("FATAL", "can't register test service handler: %v", err)
	// }
	// err = speechpb.RegisterSpeechHandlerFromEndpoint(ctx, mux, "127.0.0.1:"+serviceInfo.GrpcPort, opts)
	// if err != nil {
	// 	shlog.Logf("FATAL", "can't register test service handler: %v", err)
	// }

	// Start REST server
	shlog.Logf("INFO", "Start gRPC-Gateway server(TLS: %t) listening at %s", serviceInfo.TlsEnabled, serviceInfo.RestPort)
	if serviceInfo.TlsEnabled {
		err = http.ListenAndServeTLS("0.0.0.0:"+serviceInfo.RestPort, serviceInfo.CrtFilePath, serviceInfo.KeyFilePath, mux)
	} else {
		err = http.ListenAndServe("0.0.0.0:"+serviceInfo.RestPort, mux)
	}
	if err != nil {
		shlog.Logf("FATAL", "runServer error: %v", err)
	}
}
