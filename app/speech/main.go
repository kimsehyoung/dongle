package main

import (
	"os"

	"github.com/kimsehyoung/dongle/app/speech/server"
	"github.com/kimsehyoung/dongle/internal/utils/env"
	"github.com/kimsehyoung/logger"
)

func main() {

	_, exists := os.LookupEnv("GOOGLE_APPLICATION_CREDENTIALS")
	if !exists {
		logger.Fatalf("Google Credential is not configured")
	}
	serviceInfo := server.ServiceInfo{
		SpeechServiceAddr: env.GetEnvString("SPEECH_SERVICE_ADDR", "0.0.0.0:10005"),
		WhisperServerAddr: env.GetEnvString("WHISPER_SERVER_ADDR", "whisper-service:8000"),
	}
	server.StartGrpcServer(serviceInfo)
}
