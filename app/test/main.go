package main

import (
	"github.com/kimsehyoung/gopackages/shlog"

	"github.com/kimsehyoung/dongle/app/test/server"
	"github.com/kimsehyoung/dongle/internal/utils/env"
)

func main() {
	shlog.InitLogger("")

	serviceInfo := server.ServiceInfo{
		TestServiceAddr: env.GetEnvString("TEST_SERVICE_ADDR", "localhost:6572"),
	}
	server.StartServer(serviceInfo)
}
