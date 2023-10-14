package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/kimsehyoung/dongle/app/video/ent/videodbgen"
	"github.com/kimsehyoung/dongle/internal/utils/env"
	"github.com/kimsehyoung/logger"
)

const (
	awsRegion = types.BucketLocationConstraintApNortheast2
	bucket    = "dongle-video"
)

// Connect to db with video schema
func openDb() (*videodbgen.Client, error) {
	// Create DB client
	db, err := sql.Open("pgx", fmt.Sprintf("%s://%s:%s@%s:%s/%s?search_path=%s",
		dialect.Postgres,
		env.GetEnvString("DONGLE_DB_USER", "postgres"),
		env.GetEnvString("DONGLE_DB_PASSWORD", "1234"),
		env.GetEnvString("DONGLE_DB_HOST", "dongle-postgres"),
		env.GetEnvString("DONGLE_DB_PORT", "5432"),
		env.GetEnvString("DONGLE_DB_NAME", "dongle"),
		env.GetEnvString("DONGLE_DB_SCHEMA", "video"),
	))
	if err != nil {
		return nil, fmt.Errorf("Failed to create dongle DB client: %v", err)
	}
	// Create Schema
	_, err = db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %[1]s", env.GetEnvString("DONGLE_DB_SCHEMA", "video")))
	if err != nil {
		return nil, fmt.Errorf("Failed to create schema: %v", err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	client := videodbgen.NewClient(videodbgen.Driver(drv))
	return client, nil
}

func initDb() error {
	// Connect to db with video schema
	client, err := openDb()
	if err != nil {
		return err
	}
	defer client.Close()

	// Create tables
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		return fmt.Errorf("Failed to create tables: %v", err)
	}
	return nil
}

func createBucket() error {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(string(awsRegion)))
	if err != nil {
		return fmt.Errorf("Failed to load DefaultConfig: %v", err)
	}
	s3Client := s3.NewFromConfig(sdkConfig)
	_, err = s3Client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket:                    aws.String(bucket),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{LocationConstraint: awsRegion},
	})
	if err != nil {
		var bne *types.BucketAlreadyOwnedByYou
		if errors.As(err, &bne) {
			return nil
		}
		return fmt.Errorf("Failed to create %s bucket: %v", bucket, err)
	}
	return nil
}

func main() {
	err := createBucket()
	if err != nil {
		logger.Fatalf("[bucket init error] %v", err)
		fmt.Println(err)
	}
	logger.Info("Bucket has been initialized successfully.")

	err = initDb()
	if err != nil {
		logger.Fatalf("[database init error] %v", err)
	}
	logger.Info("Database has been initialized successfully.")
}
