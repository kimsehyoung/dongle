package auth

import (
	"sync"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/authpb"
	"github.com/kimsehyoung/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthService struct {
	authpb.AuthServer
	AuthClient authpb.AuthClient
}

var (
	once       sync.Once
	authClient authpb.AuthClient
)

func GetAuthClient(authServiceAddr string) authpb.AuthClient {
	once.Do(func() {
		conn, err := grpc.Dial(authServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			logger.Errorf("can't dial to Auth Serivce (%s) %v", authServiceAddr, err)
		}
		authClient = authpb.NewAuthClient(conn)
	})
	return authClient
}
