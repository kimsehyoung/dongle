package main

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/kimsehyoung/dongle/internal/utils/env"
	"google.golang.org/grpc"

	"github.com/kimsehyoung/dongle/api/gen/auth"
	"github.com/kimsehyoung/gopackages/shlog"
)

type ServiceInfo struct {
	ApiGatewayAddr    string `default:"6572"`
	TlsEnabled        bool   `default:false`
	KeyFilePath       string `default:"./server.key"`
	CrtFilePath       string `default:"./server.crt"`
	AuthServiceAddr   string `default:"auth-service:6572"`
	SpeechServiceAddr string `default:"speech-service:6572"`
}

func startServer(serviceInfo ServiceInfo) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lis, err := net.Listen("tcp", serviceInfo.ApiGatewayAddr)
	if err != nil {
		shlog.Logf("FATAL", "can't run server: %v", err)
	}

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	mux := runtime.NewServeMux()
	auth.RegisterAuthHandlerFromEndpoint(ctx, mux, serviceInfo.AuthServiceAddr, opts)

	if serviceInfo.TlsEnabled {
		err = http.ServeTLS(lis, mux, serviceInfo.CrtFilePath, serviceInfo.KeyFilePath)
	} else {
		err = http.Serve(lis, mux)
	}
	if err != nil {
		shlog.Logf("FATAL", "runServer error: %v", err)
	}
}

func main() {
	shlog.InitLogger("")

	serviceInfo := ServiceInfo{
		TlsEnabled:        env.GetEnvBool("TLS_ENABLED", false),
		KeyFilePath:       env.GetEnvString("KEY_FILE_PATH", "server.key"),
		CrtFilePath:       env.GetEnvString("CRT_FILE_PATH", "server.crt"),
		ApiGatewayAddr:    env.GetEnvString("API_GATEWAY_ADDR", "localhost:6572"),
		AuthServiceAddr:   env.GetEnvString("AUTH_SERVICE_ADDR", "auth-service:6572"),
		SpeechServiceAddr: env.GetEnvString("SPEECH_SERVICE_ADDR", "speech-service:6572"),
	}
	startServer(serviceInfo)
}
