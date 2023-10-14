package main

import (
	"github.com/kimsehyoung/dongle/app/auth/server"
	"github.com/kimsehyoung/dongle/app/auth/server/authdb"
	"github.com/kimsehyoung/dongle/internal/utils/env"
)

func main() {

	serviceInfo := server.ServiceInfo{
		AuthServiceAddr: env.GetEnvString("AUTH_SERVICE_ADDR", "0.0.0.0:10004"),
		DatabaseInfo: authdb.DatabaseInfo{
			Dialect:  env.GetEnvString("DONGLE_DB_DIALECT", "postgres"),
			Host:     env.GetEnvString("DONGLE_DB_HOST", "dongle-postgres"),
			Port:     env.GetEnvString("DONGLE_DB_PORT", "5432"),
			User:     env.GetEnvString("DONGLE_DB_USER", "postgres"),
			DbName:   env.GetEnvString("DONGLE_DB_NAME", "dongle"),
			Password: env.GetEnvString("DONGLE_DB_PASSWORD", "1234"),
			Schema:   env.GetEnvString("AUTH_DB_SCHEMA", "auth"),
		},
	}
	server.StartGrpcServer(serviceInfo)
}
