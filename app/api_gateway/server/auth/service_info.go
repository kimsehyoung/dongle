package auth

import (
	"github.com/kimsehyoung/dongle/api/proto/gen/go/authpb"
	"github.com/kimsehyoung/gopackages/shlog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthService struct {
	authpb.AuthServer
	AuthClient authpb.AuthClient
}

func GetAuthClient(authServiceAddr string) authpb.AuthClient {
	conn, err := grpc.Dial(authServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		shlog.Logf("ERROR", "can't dial to Auth Serivce (%s) %v", authServiceAddr, err)
	}
	return authpb.NewAuthClient(conn)
}
