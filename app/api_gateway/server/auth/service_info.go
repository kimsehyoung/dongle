package auth

import (
	"sync"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/authpb"
	"github.com/kimsehyoung/logger"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthService struct {
	authpb.AuthServer
	AuthClient    authpb.AuthClient
	TokenDbClient *redis.Client
}

var (
	authClient authpb.AuthClient
	dbOnce     sync.Once
	// For refresh token
	tokenDbClient *redis.Client
	redisOnce     sync.Once
)

func GetAuthClient(authServiceAddr string) authpb.AuthClient {
	dbOnce.Do(func() {
		conn, err := grpc.Dial(authServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			logger.Errorf("can't dial to Auth Serivce (%s) %v", authServiceAddr, err)
		}
		authClient = authpb.NewAuthClient(conn)
	})
	return authClient
}

func GetTokenDbClient(redisAddr string) *redis.Client {
	redisOnce.Do(func() {
		tokenDbClient = redis.NewClient(&redis.Options{
			Addr:     ":6379",
			Password: "",
			DB:       0,
		})
	})
	return tokenDbClient
}
