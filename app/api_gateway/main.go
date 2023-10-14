package main

import (
	"github.com/kimsehyoung/dongle/app/api_gateway/server"
	"github.com/kimsehyoung/dongle/internal/utils/env"
)

func main() {

	serviceInfo := server.ServiceInfo{
		AllowedOrigins:    env.GetEnvString("ALLOWED_ORIGINS", "http://localhost:8080"),
		RestPort:          env.GetEnvString("API_GATEWATY_REST_PORT", "10001"),
		GrpcPort:          env.GetEnvString("API_GATEWATY_GRPC_PORT", "10002"),
		TlsEnabled:        env.GetEnvBool("TLS_ENABLED", false),
		KeyFilePath:       env.GetEnvString("KEY_FILE_PATH", "ssl.key"),
		CrtFilePath:       env.GetEnvString("CRT_FILE_PATH", "ssl.crt"),
		AuthServiceAddr:   env.GetEnvString("AUTH_SERVICE_ADDR", "localhost:10004"),
		SpeechServiceAddr: env.GetEnvString("SPEECH_SERVICE_ADDR", "localhost:10005"),
		VideoServiceAddr:  env.GetEnvString("VIDEO_SERVICE_ADDR", "localhost:10006"),
		TokenDbAddr:       env.GetEnvString("TOKEN_DB_ADDR", "token-db:6379"),
	}
	server.StartGrpcServer(serviceInfo)
	server.StartRestServer(serviceInfo)
}
