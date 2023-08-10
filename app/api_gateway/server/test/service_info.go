package test

import (
	"github.com/kimsehyoung/dongle/api/proto/gen/go/testpb"
	"github.com/kimsehyoung/gopackages/shlog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TestService struct {
	testpb.TestServer
	TestClient testpb.TestClient
}

func GetTestClient(testServiceAddr string) testpb.TestClient {
	conn, err := grpc.Dial(testServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		shlog.Logf("ERROR", "can't dial to Test Serivce (%s) %v", testServiceAddr, err)
	}
	return testpb.NewTestClient(conn)
}
