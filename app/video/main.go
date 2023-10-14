package main

import (
	"github.com/kimsehyoung/dongle/app/video/server"
	"github.com/kimsehyoung/dongle/app/video/server/videodb"
	"github.com/kimsehyoung/dongle/internal/utils/env"
)

func main() {

	serviceInfo := server.ServiceInfo{
		VideoServiceAddr: env.GetEnvString("VIDEO_SERVICE_ADDR", "0.0.0.0:10006"),
		DatabaseInfo: videodb.DatabaseInfo{
			Dialect:  env.GetEnvString("DONGLE_DB_DIALECT", "postgres"),
			Host:     env.GetEnvString("DONGLE_DB_HOST", "dongle-postgres"),
			Port:     env.GetEnvString("DONGLE_DB_PORT", "5432"),
			User:     env.GetEnvString("DONGLE_DB_USER", "postgres"),
			DbName:   env.GetEnvString("DONGLE_DB_NAME", "dongle"),
			Password: env.GetEnvString("DONGLE_DB_PASSWORD", "1234"),
			Schema:   env.GetEnvString("VIDEO_DB_SCHEMA", "video"),
		},
	}
	server.StartGrpcServer(serviceInfo)
}
