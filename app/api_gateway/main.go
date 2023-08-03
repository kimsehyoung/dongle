package main

import (
	"github.com/kimsehyoung/dongle/app/api_gateway/server"
	"github.com/kimsehyoung/dongle/internal/utils/env"
	"github.com/kimsehyoung/gopackages/shlog"
)

func main() {
	shlog.InitLogger("")

	serviceInfo := server.ServiceInfo{
		TlsEnabled:        env.GetEnvBool("TLS_ENABLED", false),
		KeyFilePath:       env.GetEnvString("KEY_FILE_PATH", "server.key"),
		CrtFilePath:       env.GetEnvString("CRT_FILE_PATH", "server.crt"),
		ApiGatewayAddr:    env.GetEnvString("API_GATEWAY_ADDR", "localhost:6572"),
		TestServiceAddr:   env.GetEnvString("TEST_SERVICE_ADDR", "test-service:6572"),
		AuthServiceAddr:   env.GetEnvString("AUTH_SERVICE_ADDR", "auth-service:6572"),
		SpeechServiceAddr: env.GetEnvString("SPEECH_SERVICE_ADDR", "speech-service:6572"),
	}
	server.StartServer(serviceInfo)
}
