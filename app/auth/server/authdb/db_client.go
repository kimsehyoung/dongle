package authdb

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/kimsehyoung/dongle/app/auth/ent/authgen"
	"github.com/kimsehyoung/logger"
)

var (
	once         sync.Once
	authdbClient *authgen.Client
)

type DatabaseInfo struct {
	Dialect  string
	Host     string
	Port     string
	User     string
	DbName   string
	Password string
}

func GetDbClient(db_info DatabaseInfo) *authgen.Client {
	once.Do(func() {
		var err error

		// Create DB client
		db, err := sql.Open("pgx", fmt.Sprintf("%s://%s:%s@%s:%s/%s",
			dialect.Postgres, db_info.User, db_info.Password, db_info.Host, db_info.Port, db_info.DbName))
		if err != nil {
			logger.Fatalf("Failed to create auth DB client: %v", err)
		}
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(20)
		db.SetConnMaxLifetime(time.Hour)
		drv := entsql.OpenDB(dialect.Postgres, db)
		authdbClient = authgen.NewClient(authgen.Driver(drv))

		exists, err := authdbClient.Account.Query().Exist(context.Background())
		if err != nil {
			logger.Fatalf("Failed to check Account table: %v", err)
		}
		if !exists {
			logger.Fatal("Account table doesn't exist. Please initilaize the shema using 'init_db.go'.")
		}
	})

	return authdbClient
}
