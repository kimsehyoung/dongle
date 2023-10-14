package videodb

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/kimsehyoung/dongle/app/video/ent/videodbgen"
	"github.com/kimsehyoung/logger"
)

var (
	once          sync.Once
	videodbClient *videodbgen.Client
)

type DatabaseInfo struct {
	Dialect  string
	Host     string
	Port     string
	User     string
	DbName   string
	Password string
	Schema   string
}

func GetDbClient(db_info DatabaseInfo) *videodbgen.Client {
	once.Do(func() {
		var err error

		// Create DB client
		db, err := sql.Open("pgx", fmt.Sprintf("%s://%s:%s@%s:%s/%s?search_path=%s",
			dialect.Postgres, db_info.User, db_info.Password, db_info.Host, db_info.Port, db_info.DbName, db_info.Schema))
		if err != nil {
			logger.Fatalf("Failed to create video DB client: %v", err)
		}
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(20)
		db.SetConnMaxLifetime(time.Hour)
		drv := entsql.OpenDB(dialect.Postgres, db)
		videodbClient = videodbgen.NewClient(videodbgen.Driver(drv))

		_, err = videodbClient.OriginalVideo.Query().Exist(context.Background())
		if err != nil {
			logger.Fatalf("%v (Please initilaize using 'init_video_service.go'.)", err)
		}
	})
	return videodbClient
}
