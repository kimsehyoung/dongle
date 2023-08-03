package server

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kimsehyoung/dongle/api/proto/gen/go/authpb"
	"github.com/kimsehyoung/dongle/api/proto/gen/go/speechpb"
	"github.com/kimsehyoung/dongle/api/proto/gen/go/testpb"
	"github.com/kimsehyoung/gopackages/shlog"
	"google.golang.org/grpc"
)

type ServiceInfo struct {
	TlsEnabled        bool
	KeyFilePath       string
	CrtFilePath       string
	ApiGatewayAddr    string
	TestServiceAddr   string
	AuthServiceAddr   string
	SpeechServiceAddr string
}

type TestServer struct {
	testpb.TestServer
}

type AuthServer struct {
	authpb.AuthServer
}

type SpeechServer struct {
	speechpb.SpeechServer
}

func StartServer(serviceInfo ServiceInfo) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Register gRPC server endpoint
	opts := []grpc.DialOption{}
	mux := runtime.NewServeMux()
	testpb.RegisterTestHandlerFromEndpoint(ctx, mux, serviceInfo.TestServiceAddr, opts)
	authpb.RegisterAuthHandlerFromEndpoint(ctx, mux, serviceInfo.AuthServiceAddr, opts)
	speechpb.RegisterSpeechHandlerFromEndpoint(ctx, mux, serviceInfo.SpeechServiceAddr, opts)

	// Start server
	shlog.Logf("INFO", "Start gRPC-Gateway server(TLS: %t) listening at %s", serviceInfo.TlsEnabled, serviceInfo.ApiGatewayAddr)
	var err error
	if serviceInfo.TlsEnabled {
		err = http.ListenAndServeTLS(serviceInfo.ApiGatewayAddr, serviceInfo.CrtFilePath, serviceInfo.KeyFilePath, mux)
	} else {
		err = http.ListenAndServe(serviceInfo.ApiGatewayAddr, mux)
	}
	if err != nil {
		shlog.Logf("FATAL", "runServer error: %v", err)
	}
}
