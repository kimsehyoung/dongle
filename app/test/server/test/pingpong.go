package test

import (
	"context"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/testpb"
	"github.com/kimsehyoung/dongle/app/api_gateway/server"
	"github.com/kimsehyoung/gopackages/shlog"
)

func (s *server.TestServer) PingPong(ctx context.Context, req *testpb.PingRequest) (*testpb.PongResponse, error) {
	shlog.Logf("INFO", "Start")
	return nil, nil
}
