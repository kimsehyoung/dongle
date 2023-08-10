package main

import (
	"github.com/kimsehyoung/dongle/app/speech/server"
	"github.com/kimsehyoung/dongle/internal/utils/env"
	"github.com/kimsehyoung/gopackages/shlog"
)

func main() {
	shlog.InitLogger("")

	serviceInfo := server.ServiceInfo{
		SpeechServiceAddr: env.GetEnvString("TEST_SERVICE_ADDR", "0.0.0.0:10005"),
	}
	server.StartGrpcServer(serviceInfo)
}
