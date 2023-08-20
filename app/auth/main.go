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
			Dialect:      env.GetEnvString("AUTH_DB_DIALECT", "postgres"),
			Host:         env.GetEnvString("AUTH_DB_HOST", "auth-postgres"),
			Port:         env.GetEnvString("AUTH_DB_PORT", "5432"),
			User:         env.GetEnvString("AUTH_DB_USER", "postgres"),
			DbName:       env.GetEnvString("AUTH_DB_NAME", "auth"),
			Password:     env.GetEnvString("AUTH_DB_PASSWORD", "1234"),
			CreateSchema: env.GetEnvBool("AUTH_DB_CREATE_SCHEMA", false),
		},
	}
	server.StartGrpcServer(serviceInfo)
}
