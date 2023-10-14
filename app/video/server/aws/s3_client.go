package aws

import (
	"context"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	*s3.Client
}

var (
	once     sync.Once
	s3Client *S3Client
)

func GetS3Client(ctx context.Context) *S3Client {
	once.Do(func() {
		sdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-2"))
		if err != nil {
			fmt.Println("Failed to load DefaultConfig")
			return
		}
		s3Client = &S3Client{s3.NewFromConfig(sdkConfig)}
	})
	return s3Client
}

// func (c *S3Client) ListObjects(ctx context.Context, bucket *string) error {
// 	out, err := c.ListObjectsV2(ctx, &s3.ListObjectsV2Input{Bucket: bucket})
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(out)
// 	return nil
// }

// func (c *S3Client) CreateObject(ctx context.Context, bucket *string, key *string, body io.Reader) error {
// 	out, err := c.PutObject(ctx, &s3.PutObjectInput{
// 		Bucket: bucket,
// 		Key:    key,
// 		Body:   body,
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(out)
// 	return nil
// }

// func (c *S3Client) DeletObject(ctx context.Context, bucket *string, key *string, body io.Reader) error {
// 	out, err := c.DeleteObject(ctx, &s3.PutObjectInput{
// 		Bucket: bucket,
// 		Key:    key,
// 		Body:   body,
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(out)
// 	return nil
// }
