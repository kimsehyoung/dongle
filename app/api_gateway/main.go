package main

import (
	"github.com/kimsehyoung/dongle/app/api_gateway/server"
	"github.com/kimsehyoung/dongle/internal/utils/env"
	"github.com/kimsehyoung/gopackages/shlog"
)

func main() {
	shlog.InitLogger("")

	serviceInfo := server.ServiceInfo{
		RestPort:          env.GetEnvString("API_GATEWATY_REST_PORT", "10001"),
		GrpcPort:          env.GetEnvString("API_GATEWATY_GRPC_PORT", "10002"),
		TlsEnabled:        env.GetEnvBool("TLS_ENABLED", false),
		KeyFilePath:       env.GetEnvString("KEY_FILE_PATH", "server.key"),
		CrtFilePath:       env.GetEnvString("CRT_FILE_PATH", "server.crt"),
		TestServiceAddr:   env.GetEnvString("TEST_SERVICE_ADDR", "test-service:10003"),
		AuthServiceAddr:   env.GetEnvString("AUTH_SERVICE_ADDR", "auth-service:10004"),
		SpeechServiceAddr: env.GetEnvString("SPEECH_SERVICE_ADDR", "speech-service:10005"),
	}
	server.StartGrpcServer(serviceInfo)
	server.StartRestServer(serviceInfo)
}
