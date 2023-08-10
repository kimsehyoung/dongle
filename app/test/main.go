package main

import (
	"github.com/kimsehyoung/gopackages/shlog"

	"github.com/kimsehyoung/dongle/app/test/server"
	"github.com/kimsehyoung/dongle/internal/utils/env"
)

func main() {
	shlog.InitLogger("")

	serviceInfo := server.ServiceInfo{
		TestServiceAddr: env.GetEnvString("TEST_SERVICE_ADDR", "0.0.0.0:10003"),
	}
	server.StartGrpcServer(serviceInfo)
}
